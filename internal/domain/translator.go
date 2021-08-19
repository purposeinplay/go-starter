package domain

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

func NewTranslator(validate *validator.Validate) ut.Translator {
	locale := "en"
	english := en.New()
	uni := ut.New(english)

	trans, _ := uni.GetTranslator(locale)
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	return trans
}
