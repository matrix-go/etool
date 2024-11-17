package dao

import (
	"context"
	"github.com/matrix-go/etool/internal/dao/account_dao"
	"github.com/matrix-go/etool/internal/dao/user_dao"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserDao_InsertUserWithTx(t *testing.T) {

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DSN: "file::memory:?cache=shared",
	}))
	require.NoError(t, err)
	err = db.AutoMigrate(&user_dao.UserModel{}, &account_dao.AccountModel{})
	require.NoError(t, err)
	udao := user_dao.New(db)
	adao := account_dao.New(db)

	now := time.Now().UTC()
	user := &user_dao.UserModel{
		Name:  "test",
		Ctime: now,
		Utime: now,
	}

	err = udao.Transaction(context.Background(), func(ctx context.Context) error {
		if insertUserErr := udao.InsertUser(ctx, user); insertUserErr != nil {
			return insertUserErr
		}
		nowTime := time.Now()
		account := &account_dao.AccountModel{
			Uid:     user.Id,
			Balance: decimal.Decimal{},
			Ctime:   nowTime,
			Utime:   nowTime,
		}
		if insertAccountErr := adao.InsertAccount(ctx, account); insertAccountErr != nil {
			return insertAccountErr
		}
		//panic("implement me")
		return nil
	})
	userByID, err := udao.GetUserByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, user, userByID)
	t.Logf("user is %+v\n", user)
	//require.NoError(t, err)

	//tx := db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//		return
	//	}
	//	if err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//	tx.Commit()
	//}()
	//ctx := context.WithValue(context.Background(), "tx", tx)
	//err = dao.InsertUser(ctx, user)
	//require.NoError(t, err)

}
