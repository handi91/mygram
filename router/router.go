package router

import (
	"fmt"
	"mygram-api/config"
	"mygram-api/controller"
	"mygram-api/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
	port := config.GetServerPortEnv()
	server := fmt.Sprintf("localhost:%s", port)

	router := gin.Default()

	user := router.Group("/users")
	user.POST("/register", c.RegisterUser)
	user.POST("/login", c.LoginUser)
	user.PUT("/:userId", middleware.Authentication(), c.UpdateUser)
	user.DELETE("/:userId", middleware.Authentication(), c.DeleteUser)

	photo := router.Group("/photos").Use(middleware.Authentication())
	photo.POST("", c.PostPhoto)
	photo.GET("", c.GetPhotos)
	photo.PUT("/:photoId", c.UpdatePhoto)
	photo.DELETE("/:photoId", c.DeletePhoto)

	return router.Run(server)
}
