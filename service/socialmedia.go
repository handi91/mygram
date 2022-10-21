package service

import (
	"errors"
	"mygram-api/models/entity"
)

func (s *Service) PostSocialMedia(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	return s.db.CreateSocialMedia(socialMedia)
}

func (s *Service) GetSocialMedias() ([]entity.SocialMedia, error) {
	return s.db.GetAllSocialmedia()
}

func (s *Service) UpdateSocialMedia(socialMediaId int, socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	userId, err := s.db.GetSocialMediaUserID(socialMediaId)
	if err != nil {
		return nil, err
	}

	if userId != socialMedia.UserID {
		return nil, errors.New("can't modify not your social media")
	}

	return s.db.UpdateSocialMedia(socialMediaId, socialMedia)
}

func (s *Service) DeleteSocialMedia(socialMediaId, userId int) error {
	usrId, err := s.db.GetSocialMediaUserID(socialMediaId)
	if err != nil {
		return err
	}

	if usrId != userId {
		return errors.New("can't delete not your social media")
	}

	return s.db.DeleteSocialMedia(socialMediaId)
}
