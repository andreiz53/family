package types

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"

	database "family/database/handlers"
	"family/utils"
	"family/validate"
)

type User struct {
	ID        pgtype.UUID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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

type UpdateUserEmailValues struct {
	Email string `validate:"required,email"`
}

type UpdateUserEmailErrors struct {
	Email string
	Other string
}

type UpdateUserPasswordValues struct {
	Password string `validate:"required,gte=8"`
}

type UpdateUserPasswordErrors struct {
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

func RegisterUserToDatabaseCreateUser(u RegisterUserValues) database.CreateUserParams {
	return database.CreateUserParams{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (values UpdateUserEmailValues) Validate() *UpdateUserEmailErrors {
	errors := &UpdateUserEmailErrors{}

	v := validate.Validator
	err := v.Struct(values)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Email":
					errors.Email = "Email is invalid"
				}
			}
		}
	}
	return errors
}

func (errors *UpdateUserEmailErrors) SetOtherError(message string) {
	errors.Other = message
}

func (values UpdateUserPasswordValues) Validate() *UpdateUserPasswordErrors {
	errors := &UpdateUserPasswordErrors{}

	v := validate.Validator
	err := v.Struct(values)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Password":
					errors.Password = "At least 8 characters"
				}
			}
		}
	}
	return errors
}

func (errors *UpdateUserPasswordErrors) SetOtherError(message string) {
	errors.Other = message
}

func DBUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		DeletedAt: utils.NullTime(dbUser.DeletedAt),
		Email:     dbUser.Email,
	}
}
