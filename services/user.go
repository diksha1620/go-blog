package services

import (
	"fmt"

	"log"

	"github.com/dish1620/helper"
	"github.com/dish1620/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidateSignup(user *models.User) (map[string]interface{}, bool) {
	temp := &models.User{}
	err := models.DB.Table("users")

	//check for errors and duplicate emails
	if user.Email != "" {
		err = err.Where("email = ?", user.Email).First(temp)
	}

	if err.RowsAffected > 0 {
		if temp.Email == user.Email && user.Email != "" && temp.Email != "" {
			return helper.Message(400, "Email Address already in use."), false
		}

	}

	return helper.Message(200, "Requirement passed"), true
}

func Signup(c *gin.Context, user *models.User) map[string]interface{} {
	// Create a new validator instance.
	validate := validator.New()
	// Validate the user struct.
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var newfield string
			if err.Field() == "FirstName" {
				newfield = "FirstName"
			} else if err.Field() == "LastName" {
				newfield = "LastName"
			} else if err.Field() == "Email" {
				newfield = "Email"
			} else if err.Field() == "UserName" {
				newfield = "UserName"
			} else {
				// Handle other validation errors if needed.
				fmt.Println("Validation Error:", newfield, err)
			}

			// Check for specific validation tags and return appropriate responses.
			if err.ActualTag() == "required" {
				response := helper.Message(400, newfield+" cannot be blank.")
				return response
			} else if err.ActualTag() == "email" {
				response := helper.Message(400, newfield+" Please enter a valid email.")
				return response
			}
		}
	}

	// Perform custom validation using the ValidateSignup function.
	if resp, ok := ValidateSignup(user); !ok {
		return resp
	}

	// Create or update the user based on the ID.
	if user.ID != 0 {
		log.Println("Update user")
		models.DB.Model(user).Where("id = ?", user.ID).Updates(models.User{
			Username: user.Username,
			Email:    user.Email,
		})

		// Find the updated user and update the 'user' variable.
		models.DB.Model(user).Where("id = ?", user.ID).Find(user)
	} else {
		log.Println("Insert user")
		// Create a new user record in the database.
		models.DB.Create(user)
	}

	// Return a success response.
	response := helper.Message(200, "User created/updated successfully.")
	return response
}
