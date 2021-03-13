package models

import (
	"github.com/pborman/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid, primaryKey, default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
