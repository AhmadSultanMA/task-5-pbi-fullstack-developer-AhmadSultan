package models

//models
import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	User      User `gorm:"foreignkey:UserID"`
	UserID    uint
	Title     string `json:"title" binding:"required"`
	Caption   string `json:"caption" binding:"required"`
	ImageLink string `json:"image_link"`
}

func (Photo) TableName() string {
	return "photo"
}
