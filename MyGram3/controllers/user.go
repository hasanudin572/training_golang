package controllers

import (
	"MyGram3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
}

// GET /Users
// Get all Users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Users []models.User
	db.Find(&Users)

	c.JSON(http.StatusOK, gin.H{"data": Users})
}

// POST /Users
// Create new User
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := models.User{Username: input.Username, Email: input.Email, Password: input.Password, Age: input.Age}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /Users/:id
// Find a User
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /Users/:id
// Update a User
func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Username = input.Username
	updatedInput.Email = input.Email
	updatedInput.Password = input.Password
	updatedInput.Age = input.Age

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /Users/:id
// Delete a User
func DeleteUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
