package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dish1620/models"
	"github.com/dish1620/services"
	"github.com/gin-gonic/gin"
)

type LoginStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var v interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	m := v.(map[string]interface{})
	email := m["email"].(string)
	password := m["password"].(string)

	resp := services.UserLogin(email, password)

	c.JSON(http.StatusOK, gin.H{"response": resp})

}

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

	if existingUser.Email != "" && existingUser.Email != updateUser.Email {
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

func CreateUser(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	returnresponse := services.Signup(c, &user)

	c.JSON(http.StatusOK, gin.H{"response": returnresponse})

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
