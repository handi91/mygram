package controller

import (
	"mygram-api/helper"
	"mygram-api/models/entity"
	"mygram-api/models/request"
	"mygram-api/models/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (c *Controller) PostComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	var req request.PostComment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error bind json request",
		})
		return
	}

	if err := helper.ValidateRequest(req, ctx); err != nil {
		return
	}

	photo, err := c.service.GetPhoto(req.PhotoID)
	if err != nil && photo == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "photo not found",
		})
		return
	}

	result, err := c.service.PostComment(&entity.Comment{
		UserID:  userId,
		PhotoID: req.PhotoID,
		Message: req.Message,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error post comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.PostComment{
		ID:        result.ID,
		Message:   result.Message,
		PhotoID:   result.PhotoID,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
	})
}

func (c *Controller) GetComments(ctx *gin.Context) {
	result, err := c.service.GetComments()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error get comments",
		})
		return
	}

	var res = make([]response.GetComments, 0)
	for _, comment := range result {
		user, err := c.service.GetUser(comment.UserID)
		var commentUser response.CommentUser
		if err == nil {
			commentUser.ID = user.ID
			commentUser.Email = user.Email
			commentUser.Username = user.Username
		}

		photo, err := c.service.GetPhoto(comment.PhotoID)
		var commentPhoto response.CommentPhoto
		if err == nil {
			commentPhoto.ID = photo.ID
			commentPhoto.Title = photo.Title
			commentPhoto.Caption = photo.Caption
			commentPhoto.PhotoUrl = photo.PhotoUrl
			commentPhoto.UserID = photo.UserID
		}
		res = append(res, response.GetComments{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User:      commentUser,
			Photo:     commentPhoto,
		},
		)
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) UpdateComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param comment id",
		})
		return
	}

	var req request.UpdateComment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error bind json request",
		})
		return
	}

	if err := helper.ValidateRequest(req, ctx); err != nil {
		return
	}

	result, err := c.service.UpdateComment(commentId, userId, req.Message)

	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "comment not found",
			})
			return
		}

		if msg == "can't modify not your comment" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error update comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateComment{
		ID:        result.ID,
		Message:   result.Message,
		PhotoID:   result.PhotoID,
		UserID:    result.UserID,
		UpdatedAt: result.UpdatedAt,
	})
}

func (c *Controller) DeleteComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param comment id",
		})
		return
	}

	err = c.service.DeleteComment(commentId, userId)
	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "comment not found",
			})
			return
		}

		if msg == "can't delete not your comment" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error delete comment",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Your comment has been successfully deleted",
	})
}
