package service

import (
	"errors"
	"mygram-api/models/entity"
)

func (s *Service) PostPhoto(photo *entity.Photo) (*entity.Photo, error) {
	return s.db.CreatePhoto(photo)
}

func (s *Service) GetPhotos() ([]entity.Photo, error) {
	return s.db.GetAllPhoto()
}

func (s *Service) GetPhoto(id int) (*entity.Photo, error) {
	return s.db.GetPhotoById(id)
}

func (s *Service) UpdatePhoto(photoId int, photo *entity.Photo) (*entity.Photo, error) {
	userId, err := s.db.GetPhotoUserID(photoId)
	if err != nil {
		return nil, err
	}

	if userId != photo.UserID {
		return nil, errors.New("can't modify not your photo")
	}

	return s.db.UpdatePhoto(photoId, photo)
}

func (s *Service) DeletePhoto(photoId, userId int) error {
	usrId, err := s.db.GetPhotoUserID(photoId)
	if err != nil {
		return err
	}

	if usrId != userId {
		return errors.New("can't delete not your photo")
	}

	return s.db.DeletePhoto(photoId)
}
