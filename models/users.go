package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `json:"nickname" validate:"nonzero"`
	Email    string `json:"email" validate:"nonzero`
	Password string `json:"password" validate:"nonzero`
}

func ValidaDadosDeAluno(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
