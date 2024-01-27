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
	UpdatePhoto(models.Photo) (models.Photo, error)
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
	return Photo, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).First(&Photo).Error
}

func (db *photoRepository) AddPhoto(userID uint, Photo models.Photo) (models.Photo, error) {
	return Photo, db.connection.Preload(clause.Associations).Where("user_id = ?", userID).Create(&Photo).Error
}

func (db *photoRepository) UpdatePhoto(Photo models.Photo) (models.Photo, error) {
	err := db.connection.Where("id=?", Photo.ID).Updates(&Photo)
	if err.Error != nil {
		return models.Photo{}, err.Error
	}
	return Photo, nil
}

func (db *photoRepository) DeletePhoto(Photo models.Photo) (models.Photo, error) {
	if err := db.connection.First(&Photo, Photo.ID).Error; err != nil {
		return Photo, err
	}
	return Photo, db.connection.Delete(&Photo).Error
}
