package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataTypeMap(t *testing.T) {
	testCases := []struct {
		name           string
		typ            string
		expectDataType string
	}{
		{
			name:           "int unsigned",
			typ:            "int unsigned",
			expectDataType: "uint32",
		},
		{
			name:           "varchar(256)",
			typ:            "varchar(256)",
			expectDataType: "string",
		},
		{
			name:           "DECIMAL(10,2)",
			typ:            "DECIMAL(10,2)",
			expectDataType: "decimal.Decimal",
		},
		{
			name:           "bigint",
			typ:            "bigint",
			expectDataType: "int64",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			colType := parseColType(tc.typ)
			if dataTypeMapFunc, exists := dataTypeMap[colType.typ]; exists {
				dataType := dataTypeMapFunc(colType)
				assert.Equal(t, tc.expectDataType, dataType)
			}
		})
	}
}
