## etool

a tool for generate code

install etool
```bash
go install github.com/matrix-go/etool
```


### 1. generate dao

generate `account_dao` and `user_dao`
```bash
etool gen dao -i tests/testdata/user.sql -o ./internal -p t
etool gen dao -i tests/testdata/account.sql -o ./internal -p t
```

```bash
go test -v tests/dao/dao_test.go
```

generate `sku_dao` with table connection
```bash
etool gen dao -d 'root:1234567@tcp(127.0.0.1:3306)/db_shop' -p t -t sku -o ./internal
```
