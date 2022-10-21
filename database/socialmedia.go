package database

import (
	"errors"
	"mygram-api/models/entity"
)

func (d *Database) CreateSocialMedia(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	err := d.db.Create(socialMedia).Error
	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (d *Database) GetAllSocialmedia() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	err := d.db.Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (d *Database) GetSocialMediaUserID(id int) (int, error) {
	var userId int
	err := d.db.Model(&entity.SocialMedia{}).Select("user_id").Where("id", id).First(&userId).Error

	return userId, err
}

// func (d *Database) GetPhotoById(id int) (*entity.Photo, error) {
// 	var photo entity.Photo
// 	err := d.db.First(&photo, "id", id).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &photo, nil
// }

func (d *Database) UpdateSocialMedia(id int, socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	query := d.db.Where("id", id).Updates(socialMedia)
	err := query.Error
	if err != nil {
		return nil, err
	}

	if err == nil && query.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	socialMedia.ID = id
	return socialMedia, nil
}

func (d *Database) DeleteSocialMedia(id int) error {
	query := d.db.Delete(&entity.SocialMedia{}, "id", id)
	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return query.Error
}
