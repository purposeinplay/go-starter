package service

import (
	"context"
	"github.com/purposeinplay/go-starter/internal/adapters/psql"
	"github.com/purposeinplay/go-starter/internal/app"
	"github.com/purposeinplay/go-starter/internal/app/command"
	"github.com/purposeinplay/go-starter/internal/app/query"
	"github.com/purposeinplay/go-starter/internal/config"
	"github.com/purposeinplay/go-starter/internal/domain/i18n"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewApplication(
	_ context.Context,
	_ *config.Config,
	db *gorm.DB,
	logger *zap.Logger,
) app.Application {
	repo := psql.NewUserRepository(db)
	validator := i18n.NewValidator()
	translation := i18n.NewTranslator(validator)

	return app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(logger, repo, validator, translation),
		},
		Queries:  app.Queries{
			FindUsers: query.NewFindUsersHandler(logger, repo),
			UserByID: query.NewUserById(logger, repo),
			UserByEmail: query.NewUserByEmail(logger, repo),
		},
	}
}