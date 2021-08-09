package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Employee struct {
	ID        int    `json:"id" validate:"required"`
	FirstName string `json:"firstname" validate:"required,length(5,20)"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (c Employee) Validate() error {
	return validation.ValidateStruct(&c,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&c.ID, validation.Required),
		// Gender is optional, and should be either "Female" or "Male".
		validation.Field(&c.FirstName, validation.Required, validation.Length(5, 20)),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&c.Lastname, validation.Required),
		// Validate Address using its own validation rules
		validation.Field(&c.Email, validation.Required, is.Email),
	)
}

func (b *Employee) TableName() string {
	return "employee"
}
