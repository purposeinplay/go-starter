package app

import (
	"github.com/purposeinplay/go-starter/internal/app/command"
	"github.com/purposeinplay/go-starter/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

type Queries struct {
	FindUsers query.FindUsersHandler
	UserByID query.UserByIdHandler
	UserByEmail query.UserByEmailHandler
}
