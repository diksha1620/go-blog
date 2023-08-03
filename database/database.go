package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dish1620/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Settings struct {
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
}

func InitializeSettings() Settings {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	switch {
	case DB_HOST == "":
		fmt.Println("Environment variable DB_HOST not set.")
		os.Exit(1)
	case DB_NAME == "":
		fmt.Println("Environment variable DB_NAME not set.")
		os.Exit(1)
	case DB_USER == "":
		fmt.Println("Environment variable DB_USER not set.")
		os.Exit(1)
	case DB_PASSWORD == "":
		fmt.Println("Environment variable DB_PASSWORD not set.")
		os.Exit(1)
	}

	settings := Settings{
		DB_HOST:     DB_HOST,
		DB_NAME:     DB_NAME,
		DB_USER:     DB_USER,
		DB_PASSWORD: DB_PASSWORD,
	}

	return settings
}

func ConnectToDatabase() {
	settings := InitializeSettings()

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", settings.DB_USER, settings.DB_PASSWORD, settings.DB_HOST, settings.DB_PORT, settings.DB_NAME)

	connection, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Migrating tables")
	//normal migration
	err = connection.Debug().AutoMigrate(&models.User{}, &models.Role{}, &models.Comment{}, &models.Post{})
	if err != nil {
		fmt.Println("error in migration: ", err)
		return
	}
	fmt.Println("Done migrating")
	fmt.Println("Successfully connected!")
	DB = connection
	DB.Logger.LogMode(logger.Silent)
	fmt.Println("Done migrating")

}
