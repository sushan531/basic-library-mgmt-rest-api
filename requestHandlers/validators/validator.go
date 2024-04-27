package validators

import (
	"gopkg.in/bluesuncorp/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the input using the CustomValidator's Struct method.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
