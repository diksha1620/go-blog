package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dish1620/helper"
	"github.com/dish1620/models"
	"github.com/dish1620/services"
	"github.com/gin-gonic/gin"
)

type LoginStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// func Login(c *gin.Context) {
// 	var input LoginStruct

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	u := models.User{}

// 	u.Username = input.Username
// 	u.Password = input.Password

// 	// token, err := models.LoginCheck(u.Username, u.Password)
// 	err := models.DB.Where("id =?", c.Param("id")).First(u).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

func UpdateUser(c *gin.Context) {

	var count int64
	var user models.User
	var existingUser models.User
	var updateUser models.User

	err := models.DB.Where("id =?", c.Param("id")).First(&existingUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doees'nt exist"})
		return
	}

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existingUser.Email != "" {
		fmt.Println("not exist")
		models.DB.Where("email = ?", updateUser.Email).First(&user).Count(&count)
		if count != 0 {
			c.JSON(404, gin.H{"error": "email linked with another user, pls try different email"})
			return
		}

	}

	models.DB.Model(&existingUser).Updates(updateUser)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.FirstName = input.FirstName
	u.Username = input.Username
	u.Password = input.Password

	// SaveUser should return an error if the user couldn't be saved.
	if _, err := u.SaveUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created!"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(400, "Invalid request"))
		return
	}

	returnresponse := services.Signup(c, &user)
	helper.Respond(c.Writer, returnresponse)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Fetch all users from the database.
	models.DB.Find(&users) //adding

	// Respond with the list of users in JSON format.
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {

	var user models.User
	err := models.DB.Model(user).Where("id = ?", c.Param("id")).First(&user).Error

	// userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	// If user not found, return a JSON response with an error message.

	// c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func DeleteUser(c *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
