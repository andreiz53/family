package types

import (
	"family/validate"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string
	Password string
}

type AuthenticatedUser struct {
	Email    string
	LoggedIn bool
}

type RegisterUserValues struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=8"`
}
type RegisterUserErrors struct {
	Email    string
	Password string
	Other    string
}

type LoginUserValues struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=8"`
}
type LoginUserErrors struct {
	Email    string
	Password string
	Other    string
}

func (values RegisterUserValues) Validate() *RegisterUserErrors {
	errors := &RegisterUserErrors{}

	v := validate.Validator
	err := v.Struct(values)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Email":
					errors.Email = "Email is invalid"
				case "Password":
					errors.Password = "At least 8 characters"
				}
			}
		}
	}
	return errors
}

func (errors *RegisterUserErrors) SetOtherError(message string) {
	errors.Other = message
}

func (values LoginUserValues) Validate() *LoginUserErrors {
	errors := &LoginUserErrors{}

	v := validate.Validator
	err := v.Struct(values)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Email":
					errors.Email = "Email is invalid"
				case "Password":
					errors.Password = "At least 8 characters"
				}
			}
		}
	}
	return errors
}

func (errors *LoginUserErrors) SetOtherError(message string) {
	errors.Other = message
}
