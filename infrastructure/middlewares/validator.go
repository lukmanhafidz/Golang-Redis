package middlewares

import "github.com/go-playground/validator/v10"

func ValidateRequest(req interface{}) error {
	validate := validator.New()

	return validate.Struct(req)
}
