package server

import (
	"fmt"
	"github.com/purposeinplay/go-starter/config"
	"net/http"
)

// ListenAndServe starts the REST API.
func ListenAndServe(cfg *config.Config, h http.Handler) error {
	location := fmt.Sprintf("%v:%v", cfg.SERVER.Host, cfg.SERVER.Port)
	return http.ListenAndServe(location, h)
}