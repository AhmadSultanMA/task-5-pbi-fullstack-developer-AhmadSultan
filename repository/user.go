package repository

//repository
import (
	"RakaminProject/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepository -> User CRUD
type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetUser(uint) (models.User, error)
	GetByEmail(string) (models.User, error)
	UpdateUser(uint, map[string]interface{}) (models.User, error)
	DeleteUser(models.User) (models.User, error)
	GetAllPhoto(uint) ([]models.Photo, error)
}

type userRepository struct {
	connection *gorm.DB
}

// NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) GetUser(id uint) (user models.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) GetByEmail(email string) (user models.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) AddUser(user models.User) (models.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) UpdateUser(id uint, updateData map[string]interface{}) (models.User, error) {
	user := models.User{}
	err := db.connection.Model(&models.User{}).Where("id=?", id).Updates(updateData).Error
	if err != nil {
		return models.User{}, err
	}
	err = db.connection.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (db *userRepository) DeleteUser(user models.User) (models.User, error) {
	if err := db.connection.Model(&models.Photo{}).Where("user_id = ?", user.ID).Delete(&models.Photo{}).Error; err != nil {
		db.connection.Rollback()
		return user, err
	}
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		db.connection.Rollback()
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetAllPhoto(userID uint) (photos []models.Photo, err error) {
	return photos, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Find(&photos).Error
}
