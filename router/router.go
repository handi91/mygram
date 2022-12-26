package router

import (
// 	"fmt"
	"mygram-api/config"
	"mygram-api/controller"
	"mygram-api/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
// 	port := config.GetServerPortEnv()
// 	server := fmt.Sprintf("localhost:%s", port)
	server := "mygram-production-55c7.up.railway.app"

	router := gin.Default()

	user := router.Group("/users")
	user.POST("/register", c.RegisterUser)
	user.POST("/login", c.LoginUser)
	user.PUT("", middleware.Authentication(), c.UpdateUser)
	user.DELETE("", middleware.Authentication(), c.DeleteUser)

	photo := router.Group("/photos").Use(middleware.Authentication())
	photo.POST("", c.PostPhoto)
	photo.GET("", c.GetPhotos)
	photo.PUT("/:photoId", c.UpdatePhoto)
	photo.DELETE("/:photoId", c.DeletePhoto)

	comment := router.Group("/comments").Use(middleware.Authentication())
	comment.POST("", c.PostComment)
	comment.GET("", c.GetComments)
	comment.PUT("/:commentId", c.UpdateComment)
	comment.DELETE("/:commentId", c.DeleteComment)

	socialMedia := router.Group("/socialmedias").Use(middleware.Authentication())
	socialMedia.POST("", c.PostSocialMedia)
	socialMedia.GET("", c.GetSocialMedia)
	socialMedia.PUT("/:socialMediaId", c.UpdateSocialMedia)
	socialMedia.DELETE("/:socialMediaId", c.DeleteSocialMedia)

	return router.Run(server)
}
