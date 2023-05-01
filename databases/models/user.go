package models

import (
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primary_key" json:"id"`
	Name     string `json:"name" validate:"nonzero, regexp=^[a-zA-Z]*$"`
	UserName string `json:"username" validate:"min=4,max=40,regexp=^[a-zA-Z]*$"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"nonzero, min=8"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.NewString()
	return nil
}

func ValidUser(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
