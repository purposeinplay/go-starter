package port

import (
	"github.com/purposeinplay/go-commons/http/render"
	"github.com/purposeinplay/go-starter/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/purposeinplay/go-starter/internal/app"
	"net/http"
)

type HttpPort struct {
	app app.Application
	config *config.Config
	db *gorm.DB
	logger *zap.Logger
}

func NewHTTPPort(
	app app.Application,
	config *config.Config,
	db *gorm.DB,
	logger *zap.Logger,
) *HttpPort {

	httpPort := &HttpPort{
		app: app,
		config: config,
		db: db,
		logger: logger,
	}

	return httpPort
}

func (p *HttpPort) Healthcheck(w http.ResponseWriter, r *http.Request) error {
	return render.SendJSON(w, http.StatusOK, map[string]string{
		"name":        "Go Starter",
		"description": "An opinionated Go starter kit built on top of Chi",
	})
}