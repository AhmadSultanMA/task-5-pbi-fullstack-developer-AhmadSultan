package initializer

//initializer
import (
	"RakaminProject/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to Migrate DB")
	}
}
