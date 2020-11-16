package server

import (
	"fmt"
	config2 "github.com/oakeshq/go-starter/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

// ListenAndServe starts the REST API.
func ListenAndServe(cfg *config2.Config, h http.Handler) error {
	location := fmt.Sprintf("%v:%v", cfg.SERVER.Host, cfg.SERVER.Port)
	logrus.Infof("API started on: %s", location)
	return http.ListenAndServe(location, h)
}