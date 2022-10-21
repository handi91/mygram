package service

import (
	"errors"
	"mygram-api/models/entity"
)

func (s *Service) PostComment(comment *entity.Comment) (*entity.Comment, error) {
	return s.db.CreateComment(comment)
}

func (s *Service) GetComments() ([]entity.Comment, error) {
	return s.db.GetAllComment()
}

func (s *Service) UpdateComment(commentId, user int, message string) (*entity.Comment, error) {
	userId, err := s.db.GetCommentUserID(commentId)
	if err != nil {
		return nil, err
	}

	if userId != user {
		return nil, errors.New("can't modify not your comment")
	}

	comment, err := s.db.GetCommentByID(commentId)
	if err != nil {
		return nil, err
	}

	comment.Message = message
	return s.db.UpdateComment(commentId, comment)
}

func (s *Service) DeleteComment(commentId, userId int) error {
	usrId, err := s.db.GetCommentUserID(commentId)
	if err != nil {
		return err
	}

	if usrId != userId {
		return errors.New("can't delete not your comment")
	}

	return s.db.DeleteComment(commentId)
}
