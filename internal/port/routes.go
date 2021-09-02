package port

import (
	"github.com/go-chi/jwtauth"
	"github.com/purposeinplay/go-commons/http/router"
	"net/http"
)

type ServerInterface interface {
	ListUsers(w http.ResponseWriter, r *http.Request) error
	FindUser(w http.ResponseWriter, r *http.Request) error
	CreateUser(w http.ResponseWriter, r *http.Request) error
	Healthcheck(w http.ResponseWriter, r *http.Request) error
}

func HandlerFromMux(si ServerInterface, r router.Router) http.Handler {
	r.Group(func(r router.Router) {
		r.Get("/healthcheck", router.HandlerErrorFunc(si.Healthcheck).ServeHTTP)
		r.Get("/users", router.HandlerErrorFunc(si.ListUsers).ServeHTTP)
		r.Post("/users", router.HandlerErrorFunc(si.CreateUser).ServeHTTP)
	})

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	r.Group(func(r router.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/user", router.HandlerErrorFunc(si.FindUser).ServeHTTP)
	})

	return r
}
