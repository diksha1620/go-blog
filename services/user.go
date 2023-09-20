package services

import (
	"fmt"
	"strings"

	"log"

	"github.com/dish1620/helper"
	"github.com/dish1620/middleware"
	"github.com/dish1620/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
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

	log.Println("Insert user")
	// Create a new user record in the database.
	models.DB.Create(user)

	// Return a success response.
	response := helper.Message(200, "User created/updated successfully.")
	return response
}

type UserLoginResponse struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty" validate:"required,email"`
}

func UserLogin(email, password string) map[string]interface{} {
	validate := validator.New()
	user := &models.User{}

	pass_errs := validate.Var(password, "required")

	if pass_errs != nil {
		response := helper.Message(400, "Password cannot be blank.")
		return response
	}
	checkEmail := strings.Contains(email, "@")
	if checkEmail == true {
		errs := validate.Var(email, "required,email")

		if errs != nil {
			for _, err := range errs.(validator.ValidationErrors) {
				if err.ActualTag() == "required" {
					response := helper.Message(1, "Email cannot be blank.")
					return response
				} else if err.ActualTag() == "email" {
					response := helper.Message(400, "Please enter a valid email address.")
					return response
				}
			}
		}
		err := models.DB.Table("users").Where("email = ?", email).First(user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return helper.Message(400, "Your email address is not registered with us. Please signup to create a new account.")
			}
			return helper.Message(400, "Connection error. Please retry.")
		}

		// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		// if err != nil && (err == bcrypt.ErrMismatchedHashAndPassword || err == bcrypt.ErrHashTooShort) { //Password does not match!
		// 	return helper.Message(400, "The password you entered is incorrect. Please check again or click on forgot password to reset your password.")
		// }

	}
	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		return helper.Message(400, "Failed to generate token")
	}

	userResponse := UserLoginResponse{
		ID:       user.ID,
		UserName: user.Username,
		Email:    user.Email,
	}

	resp := helper.Message(200, "Success")
	resp["data"] = userResponse
	resp["token"] = token
	return resp
}
