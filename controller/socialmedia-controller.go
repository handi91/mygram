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

func (c *Controller) PostSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	var req request.SocialMediaRequest
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

	result, err := c.service.PostSocialMedia(&entity.SocialMedia{
		UserID:         userId,
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error post social media",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.PostSocialMedia{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserID:         result.UserID,
		CreatedAt:      result.CreatedAt,
	})
}

func (c *Controller) GetSocialMedia(ctx *gin.Context) {
	result, err := c.service.GetSocialMedias()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error get social medias",
		})
		return
	}

	var res = make([]response.GetSocialMedia, 0)
	for _, sm := range result {
		user, err := c.service.GetUser(sm.UserID)
		var socialMediaUser response.SocialMediaUser
		if err == nil {
			socialMediaUser.ID = user.ID
			socialMediaUser.Username = user.Username
		}
		res = append(res, response.GetSocialMedia{
			ID:             sm.ID,
			Name:           sm.Name,
			SocialMediaUrl: sm.SocialMediaUrl,
			UserID:         sm.UserID,
			CreatedAt:      sm.CreatedAt,
			UpdatedAt:      sm.UpdatedAt,
			User:           socialMediaUser,
		},
		)
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) UpdateSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param social media id",
		})
		return
	}

	var req request.SocialMediaRequest
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

	result, err := c.service.UpdateSocialMedia(socialMediaId, &entity.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         userId,
	})

	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "social media not found",
			})
			return
		}

		if msg == "can't modify not your social media" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error update social media",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateSocialMedia{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserID:         result.UserID,
		UpdatedAt:      result.UpdatedAt,
	})
}

func (c *Controller) DeleteSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param social media id",
		})
		return
	}

	err = c.service.DeleteSocialMedia(socialMediaId, userId)
	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "social media not found",
			})
			return
		}

		if msg == "can't delete not your social media" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": msg,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error delete social media",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Your social media has been successfully deleted",
	})
}
