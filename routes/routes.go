package routes

import (
	"github.com/dish1620/controller"
	"github.com/dish1620/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authenticated := r.Group("/api")
	authenticated.Use(middleware.AuthMiddleware())

	authenticated.POST("/register", controller.CreateUser)
	authenticated.GET("/users", controller.GetAllUsers)
	authenticated.GET("/users/:id", controller.GetUser)
	authenticated.DELETE("/deleteusers/:id", controller.DeleteUser)
	authenticated.PATCH("/users/:id", controller.UpdateUser)

	r.Run(":8080")
	return r
}
