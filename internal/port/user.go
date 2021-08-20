package port

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/jwtauth"
	commonHttp "github.com/purposeinplay/go-commons/http"
	"github.com/purposeinplay/go-commons/http/render"
	"github.com/purposeinplay/go-starter/internal/app/command"
	"github.com/purposeinplay/go-starter/internal/app/query"
	"github.com/purposeinplay/go-starter/internal/domain"
	"github.com/purposeinplay/go-starter/internal/domain/user"
	"go.uber.org/zap"
	"net/http"
)

type ListUsersRes struct {
	Users *[]user.User `json:"users"`
}

func (l *ListUsersRes) MarshalJSON() ([]byte, error) {
	type UserAlias struct {
		Id    string `json:"id"`
		Email    string `json:"email"`
	}

	users := make([]UserAlias, 0)

	for _, u := range *l.Users {
		users = append(users, UserAlias{
			Id: u.ID.String(),
			Email: u.Email,
		})
	}

	return json.Marshal(&struct {
		Users []UserAlias `json:"users"`
	}{
		Users:    users,
	})
}

type FindUserRes struct {
	User *user.User
}

func (l *FindUserRes) MarshalJSON() ([]byte, error) {
	type UserAlias struct {
		Id    string `json:"id"`
		Email    string `json:"email"`
	}

	u := &UserAlias{
		Id: l.User.ID.String(),
		Email: l.User.Email,
	}

	return json.Marshal(&struct {
		User *UserAlias `json:"user"`
	}{
		User:    u,
	})
}

type CreateUserReq struct {
	Email string `json:"email"`
}

type CreateUserRes struct {
	User *user.User `json:"user"`
	Errors interface{} `json:"errors"`
}

func (l *CreateUserRes) MarshalJSON() ([]byte, error) {
	type UserAlias struct {
		Id    string `json:"id"`
		Email    string `json:"email"`
	}

	var u *UserAlias

	if l.User != nil {
		u.Id = l.User.ID.String()
		u.Email = l.User.Email
	}

	return json.Marshal(&struct {
		User *UserAlias `json:"user"`
		Errors interface{} `json:"errors"`
	}{
		User:    u,
		Errors: l.Errors,
	})
}

func (p *HttpPort) ListUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	users, err := p.app.Queries.FindUsers.Handle(ctx, query.FindUsersCmd{})

	if err != nil {
		p.logger.Error("queries.FindUsers.Handle error", zap.Error(err))
		return commonHttp.BadRequestError("An error occurred").WithInternalMessage("could not get users: %v", err)
	}

	return render.SendJSON(w, http.StatusOK, &ListUsersRes{
		Users: users,
	})
}

func (p *HttpPort) FindUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	_, claims, _ := jwtauth.FromContext(r.Context())

	u, err := p.app.Queries.UserByID.Handle(ctx, query.UserByIdCmd{Id: claims["sub"].(string)})

	if err != nil {
		p.logger.Error("queries.FindUsers.Handle error", zap.Error(err))
		return commonHttp.BadRequestError("An error occurred").WithInternalMessage("could not find user: %v", err)
	}

	return render.SendJSON(w, http.StatusOK, &FindUserRes{
		User: u,
	})
}

func (p *HttpPort) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	createUserReq := &CreateUserReq{}
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(createUserReq)

	if err != nil {
		return commonHttp.BadRequestError("An error occurred").WithInternalMessage("could not read request: %v", err)
	}

	u, err := p.app.Commands.CreateUser.Handle(
		ctx,
		command.CreateUserCmd{
			Email:          createUserReq.Email,
		},
	)

	if err != nil {
		p.logger.Error("Commands.CreateUser error", zap.Error(err))

		var errCommon *user.Error
		var errValidation domain.ValidationError

		switch {
		case errors.As(err, &errValidation):
			return render.SendJSON(w, http.StatusUnprocessableEntity, &CreateUserRes{
				Errors: err,
			})
		case errors.As(err, &errCommon):
			// handle a specific error returned by app layer
			return commonHttp.BadRequestError("could not create user").WithInternalMessage("could not create user: %v", err)
		}

		return commonHttp.InternalServerError("An error occurred").WithInternalMessage("could not create user: %v", err)
		}

		return render.SendJSON(w, http.StatusOK, &CreateUserRes{
			User: u,
		})
}