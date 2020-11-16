package healthcheck

import (
	"github.com/oakeshq/go-starter/pkg/render"
	"github.com/oakeshq/go-starter/pkg/router"
	"net/http"
)

func RegisterHandlers(r *router.Router) {
	r.Get("/health", HealthCheck)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return render.SendJSON(w, http.StatusOK, map[string]string{
		"name":        "Go Starter",
		"description": "An opinionated Go starter kit built on top of Chi",
	})
}