package domain

import (
	"fmt"
	"time"

	"github.com/pborman/uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type ValidationError struct {
	s string `json:"s"`
	Errors map[string][]string `json:"errors"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", v.s)
}