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

	return router.Run(server)
}
