package main

import (
	"github.com/dish1620/database"
	"github.com/dish1620/models"
)

func init() {
	database.InitializeSettings()
	database.ConnectToDatabase()
}

func main() {
	database.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Comment{},
		&models.Post{},
	)
}
