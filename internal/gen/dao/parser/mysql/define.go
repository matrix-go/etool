package mysql

import "strings"

/*
	gorm.io/gen/internal/model/base.go datatype
*/

type columnType struct {
	typ      string
	realType string
	unsigned bool
}

func parseColType(colType string) columnType {
	colType = strings.ToLower(strings.TrimSpace(colType))
	splits := strings.Split(colType, " ")
	if len(splits) >= 2 {
		realType := splits[0]
		typ := realType
		idx := strings.Index(splits[0], "(")
		if idx > 0 {
			typ = splits[0][:idx]
		}
		unsigned := false
		if strings.Contains(colType, "unsigned") {
			unsigned = true
		}
		return columnType{
			typ:      typ,
			realType: realType,
			unsigned: unsigned,
		}
	} else if len(splits) == 1 {
		realType := splits[0]
		typ := realType
		idx := strings.Index(splits[0], "(")
		if idx > 0 {
			typ = splits[0][:idx]
		}
		return columnType{
			typ:      typ,
			realType: realType,
			unsigned: false,
		}
	} else {
		return columnType{
			typ:      colType,
			realType: colType,
			unsigned: false,
		}
	}
}

type dataTypeMapping func(columnType) string

var dataTypeMap = map[string]dataTypeMapping{
	"numeric": func(colType columnType) string {
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
	"integer": func(colType columnType) string {
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
	"int": func(colType columnType) string {
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
	"smallint": func(colType columnType) string {
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
	"mediumint": func(colType columnType) string {
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
	"bigint": func(colType columnType) string {
		if colType.unsigned {
			return "uint64"
		}
		return "int64"
	},
	"float":      func(colType columnType) string { return "float32" },
	"real":       func(colType columnType) string { return "float64" },
	"double":     func(colType columnType) string { return "float64" },
	"decimal":    func(colType columnType) string { return "decimal.Decimal" },
	"char":       func(colType columnType) string { return "string" },
	"varchar":    func(colType columnType) string { return "string" },
	"tinytext":   func(colType columnType) string { return "string" },
	"mediumtext": func(colType columnType) string { return "string" },
	"longtext":   func(colType columnType) string { return "string" },
	"binary":     func(colType columnType) string { return "[]byte" },
	"varbinary":  func(colType columnType) string { return "[]byte" },
	"tinyblob":   func(colType columnType) string { return "[]byte" },
	"blob":       func(colType columnType) string { return "[]byte" },
	"mediumblob": func(colType columnType) string { return "[]byte" },
	"longblob":   func(colType columnType) string { return "[]byte" },
	"text":       func(colType columnType) string { return "string" },
	"json":       func(colType columnType) string { return "string" },
	"enum":       func(colType columnType) string { return "string" },
	"time":       func(colType columnType) string { return "time.Time" },
	"date":       func(colType columnType) string { return "time.Time" },
	"datetime":   func(colType columnType) string { return "time.Time" },
	"timestamp":  func(colType columnType) string { return "time.Time" },
	"year":       func(colType columnType) string { return "int32" },
	"bit":        func(colType columnType) string { return "[]uint8" },
	"boolean":    func(colType columnType) string { return "bool" },
	"tinyint": func(colType columnType) string {
		if strings.HasPrefix(strings.TrimSpace(colType.realType), "tinyint(1)") {
			return "bool"
		}
		if colType.unsigned {
			return "uint32"
		}
		return "int32"
	},
}
