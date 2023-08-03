package main

import (
	"github.com/dish1620/controller"
	"github.com/dish1620/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {

	// // initializers.LoadEnvVariables()
	database.ConnectToDatabase()
	// // initializers.InitializeSettings()
}

func main() {
	r := gin.Default()
	r.POST("/createuser", controller.CreateUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetUser)
	r.DELETE("/deleteusers/:id", controller.DeleteUser)

	r.Run(":8080")
}
