package repository

//repository
import (
	"RakaminProject/models"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func DB() *gorm.DB {

	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Error")
	}

	db.AutoMigrate(&models.User{}, &models.Photo{})
	return db

}
