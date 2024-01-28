package repository

//repository
import (
	"RakaminProject/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository interface {
	GetPhoto(uint, int) (models.Photo, error)
	AddPhoto(uint, models.Photo) (models.Photo, error)
	UpdatePhoto(uint, map[string]interface{}) (models.Photo, error)
	DeletePhoto(models.Photo) (models.Photo, error)
}

type photoRepository struct {
	connection *gorm.DB
}

func NewPhotoRepository() PhotoRepository {
	return &photoRepository{
		connection: DB(),
	}
}

func (db *photoRepository) GetPhoto(userID uint, id int) (Photo models.Photo, err error) {
	return Photo, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Where("id = ?", id).First(&Photo).Error
}

func (db *photoRepository) AddPhoto(userID uint, Photo models.Photo) (models.Photo, error) {
	return Photo, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Create(&Photo).Error
}

func (db *photoRepository) UpdatePhoto(id uint, updateData map[string]interface{}) (models.Photo, error) {
	photo := models.Photo{}

	err := db.connection.Model(&models.Photo{}).Where("id=?", id).Updates(updateData).Error
	if err != nil {
		return models.Photo{}, err
	}
	err = db.connection.First(&photo, id).Error
	if err != nil {
		return models.Photo{}, err
	}

	return photo, nil
}

func (db *photoRepository) DeletePhoto(Photo models.Photo) (models.Photo, error) {
	if err := db.connection.First(&Photo, Photo.ID).Error; err != nil {
		return Photo, err
	}
	return Photo, db.connection.Delete(&Photo).Error
}
