package mysql

import (
	"bytes"
	"fmt"
	"github.com/xwb1989/sqlparser"
	"strings"
	"unicode"
)

type MetadataParser struct {
	ddl    *sqlparser.DDL
	prefix string
}

func NewMetadataParser(ddl *sqlparser.DDL, prefix string) *MetadataParser {
	return &MetadataParser{
		ddl:    ddl,
		prefix: prefix,
	}
}

func (p *MetadataParser) ParsePackageName() string {
	packageName := p.ddl.NewName.Name.String()
	packageName = strings.TrimPrefix(packageName, p.prefix+"_")
	packageName = strings.ToLower(packageName) + "_dao"
	return packageName
}

func (p *MetadataParser) ParseTableName() string {
	tableName := p.ddl.NewName.Name.String()
	tableName = strings.TrimPrefix(tableName, p.prefix+"_")
	tableName = p.snakeToUpperCamelCase(tableName)
	return tableName
}

func (p *MetadataParser) ParseInstanceName() string {

	tableName := p.ddl.NewName.Name.String()
	tableName = strings.TrimPrefix(tableName, p.prefix+"_")
	tableName = p.snakeToLowerCamelCase(tableName)
	return tableName
}

func (p *MetadataParser) ParseOriginTableName() string {
	return p.ddl.NewName.Name.String()
}

func (p *MetadataParser) ParseColumns() ([]Column, bool) {
	hasDecimal := false
	columns := make([]Column, 0, len(p.ddl.TableSpec.Columns))
	for _, col := range p.ddl.TableSpec.Columns {
		// TODO: special like varchar(100)
		// import should change
		colType, hasDec := p.parseColumnType(col)
		if hasDec {
			hasDecimal = true
		}
		columns = append(columns, Column{
			Name: p.parseColumnName(col),
			Type: colType,
			Tag:  p.parseColumnTag(col),
		})
	}
	return columns, hasDecimal
}

func (p *MetadataParser) parseColumnName(col *sqlparser.ColumnDefinition) string {
	return p.snakeToUpperCamelCase(col.Name.String())
}

func (p *MetadataParser) parseColumnType(col *sqlparser.ColumnDefinition) (string, bool) {
	colType := parseColType(col.Type.Type)
	hasDecimal := colType.typ == "decimal"
	if typ, exists := dataTypeMap[colType.typ]; exists {
		return typ(colType), hasDecimal
	}
	// TODO: maybe more actions
	return "any", hasDecimal
}

func (p *MetadataParser) parseColumnTag(col *sqlparser.ColumnDefinition) string {
	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s\"`", col.Name, p.toSnake(col.Name.String()))
}

func (p *MetadataParser) snakeToUpperCamelCase(name string) string {
	splits := strings.Split(name, "_")
	var sb strings.Builder
	for _, split := range splits {
		length := len(split)
		if length == 1 {
			sb.Write(bytes.ToUpper([]byte{split[0]}))
		} else if length > 1 {
			sb.Write(bytes.ToUpper([]byte{split[0]}))
			sb.WriteString(strings.ToLower(split[1:]))
		}
	}
	return sb.String()
}

func (p *MetadataParser) snakeToLowerCamelCase(name string) string {
	splits := strings.Split(name, "_")
	var sb strings.Builder
	for _, split := range splits {
		length := len(split)
		if length == 1 {
			sb.Write(bytes.ToLower([]byte{split[0]}))
		} else if length > 1 {
			if sb.Len() == 0 {
				sb.Write(bytes.ToLower([]byte{split[0]}))
			} else {
				sb.Write(bytes.ToUpper([]byte{split[0]}))
			}
			sb.WriteString(strings.ToLower(split[1:]))
		}
	}
	return sb.String()
}

func (p *MetadataParser) toSnake(name string) string {
	var sb strings.Builder
	nameRunes := []rune(name)
	for idx, r := range nameRunes {
		if unicode.IsUpper(r) {
			if sb.Len() > 0 && idx > 0 && unicode.IsLower(nameRunes[idx-1]) {
				sb.WriteString("_")
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (p *MetadataParser) ParseMethodReceiver() string {
	receiver := ""
	tableName := p.ParseTableName()
	if len(tableName) > 0 {
		receiver = strings.ToLower(string(tableName[0]))
	} else {
		receiver = "d"
	}
	return receiver
}
