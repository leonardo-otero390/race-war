package models

import (
	"net/mail"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"size:255;not null;unique" json:"nickname" validate:"nonzero"`
	Email    string `gorm:"size:100;not null;unique" json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"min=6"`
}

func (user *User) validateEmail() error {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}
	return nil
}

func ValidateUser(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	if err := user.validateEmail(); err != nil {
		return err
	}

	return nil
}
