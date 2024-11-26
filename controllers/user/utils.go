package user

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func validateInput(input InputRegisterUser) []string {
	var errors []string
	if err := validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			errors = append(errors, formatValidationError(field, err.Tag()))
		}
	}
	return errors
}

func formatValidationError(field, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return "invalid email format"
	case "min":
		return field + " is too short"
	case "max":
		return field + " is too long"
	default:
		return "invalid value for " + field
	}
}
