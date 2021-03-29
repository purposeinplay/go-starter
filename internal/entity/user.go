package entity

import (
	"github.com/purposeinplay/go-starter/pkg/storage"
)

// User model
type User struct {
	storage.Base

	Email string 				`json:"email" gorm:"uniqueIndex; not null"`
}