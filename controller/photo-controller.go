package controller

import (
	"mygram-api/models/entity"
	"mygram-api/models/request"
	"mygram-api/models/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (c *Controller) PostPhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	var req request.PhotoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error bind json request",
		})
		return
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	result, err := c.service.PostPhoto(&entity.Photo{
		UserID:   userId,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error post photo",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.PostPhoto{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		PhotoUrl:  result.PhotoUrl,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
	})
}

func (c *Controller) GetPhotos(ctx *gin.Context) {
	result, err := c.service.GetPhotos()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error get photos",
		})
		return
	}

	var res = make([]response.GetPhoto, 0)
	for _, p := range result {
		user, err := c.service.GetUser(p.UserID)
		if err != nil {
			continue
		}
		res = append(res, response.GetPhoto{
			ID:        p.ID,
			UserID:    p.UserID,
			Title:     p.Title,
			Caption:   p.Caption,
			PhotoUrl:  p.PhotoUrl,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			User: response.PhotoUser{
				Email:    user.Email,
				Username: user.Username,
			},
		},
		)
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) UpdatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param photo id",
		})
		return
	}

	var req request.PhotoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error bind json request",
		})
		return
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	result, err := c.service.UpdatePhoto(photoId, &entity.Photo{
		UserID:   userId,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	})

	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "photo not found",
			})
			return
		}

		if msg == "can't modify not your own photo" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error update photo",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdatePhoto{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		PhotoUrl:  result.PhotoUrl,
		UserID:    result.UserID,
		UpdatedAt: result.UpdatedAt,
	})
}

func (c *Controller) DeletePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param photo id",
		})
		return
	}

	err = c.service.DeletePhoto(photoId, userId)
	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "photo not found",
			})
			return
		}

		if msg == "can't delete not your own photo" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error delete photo",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Your photo has been successfully deleted",
	})
}
