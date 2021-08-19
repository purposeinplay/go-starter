package app

import (
	"context"
	"github.com/purposeinplay/go-starter/internal/app/command"
	"github.com/purposeinplay/go-starter/internal/app/query"
	"github.com/purposeinplay/go-starter/internal/config"
	"github.com/purposeinplay/go-starter/internal/domain"
	"github.com/purposeinplay/go-starter/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func NewApplication(
	_ context.Context,
	_ *config.Config,
	db *gorm.DB,
	logger *zap.Logger,
) Application {
	repo := repository.NewUserRepository(db)
	validator := domain.NewValidator()
	translation := domain.NewTranslator(validator)

	return Application{
		Commands: Commands{
			CreateUser: command.NewCreateUserHandler(logger, repo, validator, translation),
		},
		Queries:  Queries{
			FindUsers: query.NewFindUsersHandler(logger, repo),
			UserByID: query.NewUserById(logger, repo),
			UserByEmail: query.NewUserByEmail(logger, repo),
		},
	}
}