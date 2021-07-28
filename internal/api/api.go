package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	mw "github.com/purposeinplay/go-commons/http/middleware"
	"github.com/purposeinplay/go-commons/logs"
	"github.com/purposeinplay/go-starter/config"
	"github.com/purposeinplay/go-starter/internal/repository"
	service2 "github.com/purposeinplay/go-starter/internal/service"
	"github.com/rs/cors"
	"gorm.io/gorm"

	"github.com/purposeinplay/go-commons/http/router"
)

// API exposes the integral struct
type API struct {
	handler http.Handler
	r       *router.Router
	config  *config.Config
	db      *gorm.DB
	service *service2.Service
}

// NewAPI instantiates a new REST API.
func NewAPI(
	config *config.Config,
	r *router.Router,
	db *gorm.DB,
) *API {
	logger := logs.NewLogger()
	defer logger.Sync()

	api := &API{
		r:      r,
		config: config,
		db:     db,
	}

	repo := repository.NewRepository(db)
	service := service2.NewService(repo)
	api.service = service
	ctx := context.Background()
	r.Chi.Use(middleware.RealIP)
	r.Chi.Use(middleware.RequestID)
	r.Chi.Use(mw.NewLoggerMiddleware(logger))
	r.Use(mw.Recoverer)

	corsHandler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link", "X-Total-Count"},
		AllowCredentials: true,
	})

	r.Route("/v1", func(r *router.Router) {
		r.Get("/users", api.ListUsers)
	})

	r.Get("/health", HealthCheck)

	api.handler = corsHandler.Handler(chi.ServerBaseContext(ctx, r))
	return api
}
