package models

//models
import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" `
	NoTelp   string `json:"notelp"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !isValidPassword(u.Password) {
		err = errors.New("password must be at least 6 words")
	}
	return
}

func isValidPassword(password string) bool {
	return len(password) >= 6
}

func (User) TableName() string {
	return "users"
}
