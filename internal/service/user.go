package service

import (
	"github.com/pborman/uuid"
	"github.com/purposeinplay/go-starter/internal/entity"
	"github.com/purposeinplay/go-starter/pkg/storage"
)

type repository interface {
	Find(conds ...interface{}) (*[]entity.User, error)
	First(conds ...interface{}) (*entity.User, error)
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
func (s *Service) FirstByUserID(userID uuid.UUID) (*[]entity.User, error) {
	conds := entity.User{
		Base:   storage.Base{
			ID: userID,
		},
	}

	return s.repo.Find(conds)
}

// FindAll Returns all users
func (s *Service) FindAll() (*[]entity.User, error) {
	return s.repo.Find()
}
