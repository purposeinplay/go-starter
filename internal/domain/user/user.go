package user

import (
	"github.com/purposeinplay/go-starter/internal/domain"
)

// User model
type User struct {
	domain.Base
	Email string 				`json:"email" gorm:"uniqueIndex; not null"`
}

type Error struct {
	s string
}

func (e *Error) Error() string {
	return e.s
}

func New(text string) error {
	return &Error{text}
}

var (
	ErrUserNotFound   = New("user not found")
	ErrCouldNotCreateUser   = New("could not create user")
)