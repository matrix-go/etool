package mysql

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestParser_ParseDao(t *testing.T) {
	file, err := os.OpenFile("testdata/account.sql", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	require.NoError(t, err)
	sql := string(bytes)
	p, err := NewParser(sql, "t")
	require.NoError(t, err)
	target, err := p.ParseDaoTypes()
	require.NoError(t, err)
	t.Log(target)
}

func TestParser_ParseDaoImpl(t *testing.T) {
	file, err := os.OpenFile("testdata/user.sql", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	require.NoError(t, err)
	sql := string(bytes)
	p, err := NewParser(sql, "t")
	require.NoError(t, err)
	target, err := p.ParseDaoImpl()
	require.NoError(t, err)
	t.Log(target)
}
