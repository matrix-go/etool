## etool

a tool for generate code

install etool
```bash
go install .
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
