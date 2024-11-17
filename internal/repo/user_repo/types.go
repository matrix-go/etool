package user_repo

import "context"

type UserRepo interface {
	CreateUser(ctx context.Context, user *UserEntity) error
}

type UserEntity struct {
}
