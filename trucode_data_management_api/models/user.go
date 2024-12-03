package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Username   string `json:"username" gorm:"not null;unique"`
	Password   string `json:"-"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHashed)
	return nil
}
