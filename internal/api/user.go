package api

import (
	"encoding/json"
	"net/http"

	"github.com/purposeinplay/go-commons/http/httperr"
	"github.com/purposeinplay/go-commons/http/render"
	"github.com/purposeinplay/go-starter/internal/entity"
)

type ListUsersResponse struct {
	Users *[]entity.User `json:"users"`
}

func (p *ListUsersResponse) MarshalJSON() ([]byte, error) {
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


	return render.SendJSON(w, http.StatusOK, &ListUsersResponse{
		Users: users,
	})
}
