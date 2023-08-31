package main

import (
	"os"

	"github.com/dish1620/models"
	Router "github.com/dish1620/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {

	// // initializers.LoadEnvVariables()
	models.ConnectToDatabase()
	// // initializers.InitializeSettings()
}

func main() {
	godotenv.Load()            // Load env variables
	models.ConnectToDatabase() // load db
	// We want to get the router in async, thus a channel is required to return the router instance.

	// var router = make(chan *gin.Engine)
	// go routes.GetRouter(router)
	// var port string = os.Getenv("SERVER_PORT")
	// server_addr := fmt.Sprintf(":%s", port)
	// r := <-router

	router := Router.SetupRouter()

	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
