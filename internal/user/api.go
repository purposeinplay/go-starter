package user

import (
	"github.com/oakeshq/go-starter/config"
	"github.com/oakeshq/go-starter/pkg/router"
	"gorm.io/gorm"
)

func RegisterHandlers(r *router.Router, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	r.Get("/users", handler.ListUsers)
}
