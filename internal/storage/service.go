package storage

import (
	"github.com/oakeshq/go-starter/pkg/storage/models"
	"github.com/pborman/uuid"
)

type repository interface {
	Find(conds ...interface{}) (*[]User, error)
	First(conds ...interface{}) (*User, error)
}

type Service struct{
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

// FirstByUserID Returns the first storage that matches
func (s *Service) FirstByUserID(userID uuid.UUID) (*[]User, error) {
	conds := User{
		Base:   models.Base{
			ID: userID,
		},
	}

	return s.repo.Find(conds)
}
// FindAll Returns all users
func (s *Service) FindAll() (*[]User, error) {
	return s.repo.Find()
}
