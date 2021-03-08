package api

import (
"encoding/json"
"github.com/oakeshq/go-starter/internal/storage"
"github.com/oakeshq/go-starter/pkg/httperr"
"github.com/oakeshq/go-starter/pkg/render"
"net/http"
)

type responsePayload struct {
	Users *[]storage.User `json:"users"`
}

func (p *responsePayload) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Email    string `json:"email"`
	}

	users := make([]Alias, 0)
	userList := *p.Users

	for _, user := range userList {
		users = append(users, Alias{
			Email: user.Email,
		})
	}

	return json.Marshal(users)
}


func (a *API) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := a.service.FindAll()

	if err != nil {
		return httperr.BadRequestError("An error occurred").WithInternalMessage("Couldn't fetch accounts: %v", err)
	}

	return render.SendJSON(w, http.StatusOK, &responsePayload{
		Users: users,
	})
}
