package models

//models
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password" `
	NoTelp   string  `json:"notelp"`
	Photos   []Photo `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}
