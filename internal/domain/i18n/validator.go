package i18n

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *validator.Validate {
	validate := validator.New()

	return validate
}