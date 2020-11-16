package user

import (
	"github.com/oakeshq/go-starter/internal/user/storage"
	"github.com/oakeshq/go-starter/pkg/storage/models"
	"github.com/pborman/uuid"
)

type repository interface {
	Find(conds ...interface{}) (*[]storage.User, error)
	First(conds ...interface{}) (*storage.User, error)
}

type Service struct{
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

// FirstByUserID Returns the first user that matches
func (s *Service) FirstByUserID(userID uuid.UUID) (*[]storage.User, error) {
	conds := storage.User{
		Base:   models.Base{
			ID: userID,
		},
	}

	return s.repo.Find(conds)
}
// FindAll Returns all users
func (s *Service) FindAll() (*[]storage.User, error) {
	return s.repo.Find()
}
