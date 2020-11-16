package user

import (
	"github.com/oakeshq/go-starter/internal/user/storage"
	"gorm.io/gorm"
)

// Repository ...
type Repository struct {
	db *gorm.DB
}

// NewRepository Returns a new instance of the payment repo
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Find all users that match the conditions
func (r *Repository) Find(conds ...interface{}) (*[]storage.User, error) {
	var users []storage.User

	if err := r.db.Find(&users, conds...).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// Return first user that matches the conditions
func (r *Repository) First(conds ...interface{}) (*storage.User, error) {
	var user storage.User

	if err := r.db.First(&user, conds...).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
