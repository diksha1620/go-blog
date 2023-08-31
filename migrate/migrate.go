package main

import (
	"github.com/dish1620/models"
)

func init() {
	models.InitializeSettings()
	models.ConnectToDatabase()
}

func main() {
	models.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Comment{},
		&models.Post{},
	)
}
