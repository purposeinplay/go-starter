package psql

import (
	"context"
	"github.com/purposeinplay/go-starter/internal/domain/user"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser ...
func (r *UserRepository) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindUser ...
func (r *UserRepository) FindUser(ctx context.Context, conds ...interface{}) (*[]user.User, error) {
	var users []user.User

	if err := r.db.Find(&users, conds...).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// FirstUser ...
func (r *UserRepository) FirstUser(ctx context.Context, conds ...interface{}) (*user.User, error) {
	var user user.User

	if err := r.db.First(&user, conds...).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
