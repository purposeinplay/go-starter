package query

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/purposeinplay/go-starter/internal/domain/user"
)

type UserByEmailCmd struct {
	Email string
}

type UserByEmailHandler struct {
	logger *zap.Logger
	repo   user.UserRepository
}

func NewUserByEmail(
	logger     *zap.Logger,
	repo user.UserRepository,
) UserByEmailHandler {
	return UserByEmailHandler{
		logger: logger,
		repo: repo,
	}
}

func (s *UserByEmailHandler) Handle(ctx context.Context, q UserByEmailCmd) (*user.User, error) {
	t, err := s.repo.First(ctx, user.User{
		Email: q.Email,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, user.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return t, nil
}
