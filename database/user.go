package database

import (
	"errors"
	"mygram-api/models/entity"
)

func (d *Database) CreateUser(userData *entity.User) (*entity.User, error) {
	err := d.db.Create(userData).Error
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (d *Database) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := d.db.First(&user, "email", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) GetUserById(id int) (*entity.User, error) {
	var user entity.User
	err := d.db.First(&user, "id", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) UpdateUser(id int, user *entity.User) error {
	query := d.db.Where("id", id).Updates(user)
	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return query.Error
}

func (d *Database) DeleteUser(id int) error {
	query := d.db.Delete(&entity.User{}, "id", id)
	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return query.Error
}
