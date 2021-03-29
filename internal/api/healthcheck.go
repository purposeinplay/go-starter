package api

import (
	"github.com/purposeinplay/go-commons/http/render"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return render.SendJSON(w, http.StatusOK, map[string]string{
		"name":        "Go Starter",
		"description": "An opinionated Go starter kit built on top of Chi",
	})
}