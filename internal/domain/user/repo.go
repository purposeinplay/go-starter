package user

import "context"

type UserRepository interface {
	FindUser(ctx context.Context, conds ...interface{}) (*[]User, error)
	FirstUser(ctx context.Context, conds ...interface{}) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
}