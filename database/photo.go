package database

import (
	"errors"
	"mygram-api/models/entity"
)

func (d *Database) CreatePhoto(photoData *entity.Photo) (*entity.Photo, error) {
	err := d.db.Create(photoData).Error
	if err != nil {
		return nil, err
	}

	return photoData, nil
}

func (d *Database) GetAllPhoto() ([]entity.Photo, error) {
	var photos []entity.Photo
	err := d.db.Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (d *Database) GetUserID(id int) (int, error) {
	var userId int
	err := d.db.Model(&entity.Photo{}).Select("user_id").Where("id", id).First(&userId).Error

	return userId, err
}

func (d *Database) UpdatePhoto(id int, photo *entity.Photo) (*entity.Photo, error) {
	query := d.db.Where("id", id).Updates(photo)
	err := query.Error
	if err != nil {
		return nil, err
	}

	if err == nil && query.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	photo.ID = id
	return photo, nil
}

func (d *Database) DeletePhoto(id int) error {
	query := d.db.Delete(&entity.Photo{}, "id", id)
	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return query.Error
}
