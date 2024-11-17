package mysql

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/matrix-go/etool/internal/gen/dao/parser/mysql/options"
	"github.com/xwb1989/sqlparser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

//go:embed template/gorm/types.tmpl
var typesTmplStr string

//go:embed template/gorm/dao.tmpl
var daoTmplStr string

type Parser interface {
	PackageName() string
	ParseDaoTypes() (types string, err error)
	ParseDaoImpl() (dao string, err error)
	Write(outPath string) error
}

type parser struct {
	ddl         *sqlparser.DDL
	typesTmpl   *template.Template
	daoTmpl     *template.Template
	typesStruct TypesParseStruct
}

func (p *parser) Write(outPath string) error {

	packageName := p.PackageName()

	// types file
	types, err := p.ParseDaoTypes()
	if err != nil {
		return err
	}

	// dao file
	dao, err := p.ParseDaoImpl()
	if err != nil {
		return err
	}

	switch outPath {
	case "stdout":
		fmt.Println(types)
	default:
		daoDir := path.Join(outPath, "dao", packageName)
		if err = os.MkdirAll(daoDir, os.ModePerm); err != nil {
			return err
		}
		typeFile := path.Join(daoDir, "types.gen.go")
		typesOutFile, err := os.OpenFile(typeFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer typesOutFile.Close()
		_, err = typesOutFile.WriteString(types)
		if err != nil {
			return err
		}
		daoFile := path.Join(daoDir, "dao.gen.go")
		daoOutFile, err := os.OpenFile(daoFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer daoOutFile.Close()
		_, err = daoOutFile.WriteString(dao)
		if err != nil {
			return err
		}

		// fmt generated code
		err = exec.Command("gofmt", "-s", "-w", daoDir).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func NewParser(opts ...options.ParserOption) (Parser, error) {

	parserOpt := &options.Option{}
	for _, opt := range opts {
		opt(parserOpt)
	}
	parserMode := parserOpt.Validate()
	switch parserMode {
	case options.ParserModeDDLPath:
		return NewDDLPathParser(parserOpt)
	case options.ParserModeDSNConnect:
		return NewDSNConnectParser(parserOpt)
	default:
		return nil, errors.New("invalid parser mode")
	}
}

func NewDDLPathParser(parserOpt *options.Option) (Parser, error) {

	file, err := os.OpenFile(parserOpt.DDLPath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	sqlBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	sql := string(sqlBytes)
	// ddl
	stmt, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return nil, err
	}
	ddl, ok := stmt.(*sqlparser.DDL)
	if !ok {
		return nil, errors.New("unsupported ddlSql")
	}
	if ddl.Action != sqlparser.CreateStr {
		return nil, errors.New("unsupported ddl, only create ddl supported")
	}

	// typesStruct
	metadataParser := NewMetadataParser(ddl, parserOpt.Prefix)
	typesParseStruct := TypesParseStruct{}
	typesParseStruct.PackageName = metadataParser.ParsePackageName()
	typesParseStruct.TableName = metadataParser.ParseTableName()
	typesParseStruct.OriginTableName = metadataParser.ParseOriginTableName()
	typesParseStruct.InstanceName = metadataParser.ParseInstanceName()
	typesParseStruct.MethodReceiver = metadataParser.ParseMethodReceiver()
	typesParseStruct.Columns, typesParseStruct.HasDecimal = metadataParser.ParseColumns()

	// dao template
	typesTemplate := template.New("typesTemplate")
	typesTmpl, err := typesTemplate.Parse(typesTmplStr)
	if err != nil {
		return nil, err
	}

	// dao impl template
	daoTemplate := template.New("daoTemplate")
	daoTmpl, err := daoTemplate.Parse(daoTmplStr)
	if err != nil {
		return nil, err
	}

	return &parser{
		ddl:         ddl,
		typesTmpl:   typesTmpl,
		daoTmpl:     daoTmpl,
		typesStruct: typesParseStruct,
	}, nil
}

func NewDSNConnectParser(parserOpt *options.Option) (Parser, error) {

	db, err := gorm.Open(mysql.Open(parserOpt.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()

	// show tables
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}

	tableExists := false
	targetTable := parserOpt.Table
	if !strings.HasPrefix(targetTable, parserOpt.Prefix) {
		targetTable = parserOpt.Prefix + "_" + targetTable
	}

	for _, table := range tables {
		if targetTable == table {
			tableExists = true
			break
		}
	}
	if !tableExists {
		return nil, errors.New("target table not found")
	}

	var table string
	var sql string
	rows, err := sqlDB.Query("SHOW CREATE TABLE " + targetTable)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&table, &sql); err != nil {
			return nil, err
		}
	}

	// ddl
	stmt, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return nil, err
	}
	ddl, ok := stmt.(*sqlparser.DDL)
	if !ok {
		return nil, errors.New("unsupported ddlSql")
	}
	if ddl.Action != sqlparser.CreateStr {
		return nil, errors.New("unsupported ddl, only create ddl supported")
	}

	// typesStruct
	metadataParser := NewMetadataParser(ddl, parserOpt.Prefix)
	typesParseStruct := TypesParseStruct{}
	typesParseStruct.PackageName = metadataParser.ParsePackageName()
	typesParseStruct.TableName = metadataParser.ParseTableName()
	typesParseStruct.OriginTableName = metadataParser.ParseOriginTableName()
	typesParseStruct.InstanceName = metadataParser.ParseInstanceName()
	typesParseStruct.MethodReceiver = metadataParser.ParseMethodReceiver()
	typesParseStruct.Columns, typesParseStruct.HasDecimal = metadataParser.ParseColumns()

	// dao template
	typesTemplate := template.New("typesTemplate")
	typesTmpl, err := typesTemplate.Parse(typesTmplStr)
	if err != nil {
		return nil, err
	}

	// dao impl template
	daoTemplate := template.New("daoTemplate")
	daoTmpl, err := daoTemplate.Parse(daoTmplStr)
	if err != nil {
		return nil, err
	}

	return &parser{
		ddl:         ddl,
		typesTmpl:   typesTmpl,
		daoTmpl:     daoTmpl,
		typesStruct: typesParseStruct,
	}, nil
}

type TypesParseStruct struct {
	PackageName     string
	TableName       string
	OriginTableName string
	InstanceName    string
	MethodReceiver  string
	HasDecimal      bool
	Columns         []Column
}

type Column struct {
	Name string
	Type string
	Tag  string
}

func (p *parser) PackageName() string {
	return p.typesStruct.PackageName
}

func (p *parser) ParseDaoTypes() (types string, err error) {
	bufferString := bytes.NewBufferString(types)
	err = p.typesTmpl.Execute(bufferString, p.typesStruct)
	return bufferString.String(), err
}

func (p *parser) ParseDaoImpl() (dao string, err error) {
	bufferString := bytes.NewBufferString(dao)
	err = p.daoTmpl.Execute(bufferString, p.typesStruct)
	return bufferString.String(), err
}
