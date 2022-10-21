package database

import (
	"errors"
	"mygram-api/models/entity"
)

func (d *Database) CreateComment(comment *entity.Comment) (*entity.Comment, error) {
	err := d.db.Create(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (d *Database) GetAllComment() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := d.db.Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (d *Database) GetCommentByID(id int) (*entity.Comment, error) {
	var comment entity.Comment
	if err := d.db.First(&comment, "id", id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (d *Database) GetCommentUserID(id int) (int, error) {
	var userId int
	err := d.db.Model(&entity.Comment{}).Select("user_id").Where("id", id).First(&userId).Error

	return userId, err
}

func (d *Database) UpdateComment(id int, comment *entity.Comment) (*entity.Comment, error) {
	query := d.db.Where("id", id).Updates(comment)
	err := query.Error
	if err != nil {
		return nil, err
	}

	if err == nil && query.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return comment, nil
}

func (d *Database) DeleteComment(id int) error {
	query := d.db.Delete(&entity.Comment{}, "id", id)
	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return query.Error
}
