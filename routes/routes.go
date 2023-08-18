package routes

import (
	"github.com/dish1620/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.POST("/createuser", controller.CreateUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetUser)
	r.DELETE("/deleteusers/:id", controller.DeleteUser)

	r.Run(":8080")
	return r
}
