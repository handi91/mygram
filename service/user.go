package service

import (
	"errors"
	"mygram-api/helper"
	"mygram-api/models/entity"
	"mygram-api/models/request"
)

func (s *Service) RegisterUser(user *entity.User) (*entity.User, error) {
	return s.db.CreateUser(user)
}

func (s *Service) LoginUser(req request.LoginUser) (string, error) {
	user, err := s.db.GetUserByEmail(req.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return "", errors.New("email not registered")
		}
		return "", err
	}

	if !helper.ComparePassword(req.Password, user.Password) {
		return "", errors.New("incorrect password")
	}

	token := helper.GenerateToken(user.ID, user.Email)
	return token, nil
}

func (s *Service) UpdateUser(id int, user *entity.User) (*entity.User, error) {
	err := s.db.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return s.db.GetUserById(id)
}

func (s *Service) DeleteUser(id int) error {
	return s.db.DeleteUser(id)
}
