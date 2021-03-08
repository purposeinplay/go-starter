package storage

import (
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
func (r *Repository) Find(conds ...interface{}) (*[]User, error) {
	var users []User

	if err := r.db.Find(&users, conds...).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// Return first storage that matches the conditions
func (r *Repository) First(conds ...interface{}) (*User, error) {
	var user User

	if err := r.db.First(&user, conds...).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
