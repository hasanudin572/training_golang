package controllers

import (
	"MyGram3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateComInput struct {
	User_id  int64  `json:"user_id"`
	Photo_id int64  `json:"photo_id"`
	Message  string `json:"message"`
}

type UpdateComInput struct {
	User_id  int64  `json:"user_id"`
	Photo_id int64  `json:"photo_id"`
	Message  string `json:"message"`
}

// GET /Coms
// Get all Coms
func FindComs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Coms []models.Comment
	db.Find(&Coms)

	c.JSON(http.StatusOK, gin.H{"data": Coms})
}

// POST /Coms
// Create new Com
func CreateCom(c *gin.Context) {
	// Validate input
	var input CreateComInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Com
	Com := models.Comment{User_id: input.User_id, Photo_id: input.Photo_id, Message: input.Message}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Com)

	c.JSON(http.StatusOK, gin.H{"data": Com})
}

// GET /Coms/:id
// Find a Com
func FindCom(c *gin.Context) { // Get model if exist
	var Com models.Comment

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Com})
}

// PATCH /Coms/:id
// Update a Com
func UpdateCom(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Com models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&Com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateComInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Comment
	updatedInput.User_id = input.User_id
	updatedInput.Photo_id = input.Photo_id
	updatedInput.Message = input.Message

	db.Model(&Com).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Com})
}

// DELETE /Coms/:id
// Delete a Com
func DeleteCom(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var com models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&com)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
