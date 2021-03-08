package storage

import (
	"github.com/oakeshq/go-starter/pkg/storage/models"
)

// User model
type User struct {
	models.Base

	Email string 				`json:"email" gorm:"uniqueIndex; not null"`
}