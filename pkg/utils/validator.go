package utils

import (
	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func ValidateStruct(s interface{}) error {

	return v.Struct(s)
}
