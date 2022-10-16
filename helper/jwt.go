package helper

import (
	"errors"
	"mygram-api/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id int, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	secretKey := config.GetSecretKeyEnv()
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseToken.SignedString([]byte(secretKey))
	return token
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	header := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(header, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(header, " ")[1]
	secretKey := config.GetSecretKeyEnv()
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
