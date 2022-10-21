package helper

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ValidateRequest(req interface{}, ctx *gin.Context) error {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return err
	}

	return nil
}
