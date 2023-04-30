package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Name     string
	UserName string
	Email    string
	Password string
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.NewString()
	return nil
}
