package controller

import (
	"mygram-api/helper"
	"mygram-api/models/entity"
	"mygram-api/models/request"
	"mygram-api/models/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var req request.RegisterUser
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

	var user = entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: helper.HashPassword(req.Password),
		Age:      req.Age,
	}

	result, err := c.service.RegisterUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "Email has already been taken",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "Username has already been taken",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error register user",
		})
		return
	}

	var res = response.RegisterUser{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		Age:      result.Age,
	}
	ctx.JSON(http.StatusCreated, res)
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	var req request.LoginUser
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

	token, err := c.service.LoginUser(req)
	if err != nil {
		msg := err.Error()
		if msg == "email not registered" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": msg,
			})
			return
		}
		if msg == "incorrect password" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": msg,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    404,
			"message": "error get user from db",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.LoginUser{Token: token})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param user id",
		})
		return
	}

	var req request.UpdateUser
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

	result, err := c.service.UpdateUser(userId, &entity.User{
		Email:    req.Email,
		Username: req.Username,
	})

	if err != nil {
		msg := err.Error()
		if msg == "user not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": msg,
			})
			return
		}

		if strings.Contains(msg, "unique") {
			if strings.Contains(msg, "email") {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "Email has already been taken",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "Username has already been taken",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateUser{
		ID:        result.ID,
		Email:     result.Email,
		Username:  result.Username,
		Age:       result.Age,
		UpdatedAt: result.UpdatedAt,
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid param user id",
		})
		return
	}

	err = c.service.DeleteUser(userId)
	if err != nil {
		msg := err.Error()
		if msg == "user not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": msg,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Your account has been successfully deleted",
	})
}
