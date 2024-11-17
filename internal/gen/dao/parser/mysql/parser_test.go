package mysql

import (
	_ "embed"
	"github.com/matrix-go/etool/internal/gen/dao/parser/mysql/options"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSqlPathParser_Parse(t *testing.T) {
	p, err := NewDDLPathParser(&options.Option{
		DDLPath: "testdata/account.sql",
		Prefix:  "t",
	})
	require.NoError(t, err)

	// types
	types, err := p.ParseDaoTypes()
	require.NoError(t, err)
	t.Log(types)

	// dao
	dao, err := p.ParseDaoImpl()
	require.NoError(t, err)
	t.Log(dao)
}

//go:embed testdata/user.sql
var userSql string

func TestConnectionParser_ParseDaoImpl(t *testing.T) {

	dsn := "root:1234567@tcp(127.0.0.1:3306)/db_shop"
	prefix := "t"

	// parser parse
	p, err := NewDSNConnectParser(&options.Option{
		DSN:    dsn,
		Prefix: prefix,
		Table:  "sku",
	})
	require.NoError(t, err)

	// types
	types, err := p.ParseDaoTypes()
	require.NoError(t, err)
	t.Log(types)

	// dao
	dao, err := p.ParseDaoImpl()
	require.NoError(t, err)
	t.Log(dao)
}
