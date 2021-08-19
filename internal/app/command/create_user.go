package command

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/purposeinplay/go-starter/internal/domain"
	"github.com/purposeinplay/go-starter/internal/domain/user"
	"go.uber.org/zap"
	"strings"
)

type CreateUserCmd struct {
	Email string `json:"email" validate:"required"`
}

type CreateUserHandler struct {
	logger            *zap.Logger
	repo              user.UserRepository
	validator *validator.Validate
	translator ut.Translator
}

func NewCreateUserHandler(
	logger     *zap.Logger,
	repo user.UserRepository,
	validator *validator.Validate,
	translator ut.Translator,
) CreateUserHandler {
	return CreateUserHandler{
		logger: logger,
		repo: repo,
		validator: validator,
		translator: translator,
	}
}

func (h *CreateUserHandler) validatorErrorsToValidationErrors(errors validator.ValidationErrors) error {
	err := domain.ValidationError{
		Errors: make(map[string][]string),
	}

	for _, e := range errors {
		field := strings.ToLower(e.Field())
		translation := e.Translate(h.translator)
		val := append(err.Errors[field], translation)
		err.Errors[field] = val
	}

	return err
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUserCmd) (*user.User, error) {
	err := h.validator.StructCtx(ctx, cmd)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}

		errors := h.validatorErrorsToValidationErrors(err.(validator.ValidationErrors))

		return nil, errors
	}

	u, err := h.repo.Create(ctx, &user.User{
		Email:             cmd.Email,
	})

	if err != nil {
		return nil, user.ErrCouldNotCreateUser
	}

	return u, nil
}