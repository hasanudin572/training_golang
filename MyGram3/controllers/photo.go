package controllers

import (
	"MyGram3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreatePhotoInput struct {
	Title   string `json:"Title"`
	Caption string `json:"Caption"`
	User_id int64  `json:"User_id"`
}

type UpdatePhotoInput struct {
	Title   string `json:"Title"`
	Caption string `json:"Caption"`
	User_id int64  `json:"User_id"`
}

// GET /photos
// Get all photos
func FindPhotos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var photos []models.Photo
	db.Find(&photos)

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

// POST /photos
// Create new photo
func CreatePhoto(c *gin.Context) {
	// Validate input
	var input CreatePhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create photo
	photo := models.Photo{Title: input.Title, Caption: input.Caption, User_id: input.User_id}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&photo)

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

// GET /photo/:id
// Find a photo
func FindPhoto(c *gin.Context) { // Get model if exist
	var photo models.Photo

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

// PATCH /photos/:id
// Update a photo
func UpdatePhoto(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var photo models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Photo
	updatedInput.Title = input.Title
	updatedInput.Caption = input.Caption
	updatedInput.User_id = input.User_id

	db.Model(&photo).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

// DELETE /photos/:id
// Delete a photo
func DeletePhoto(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var photo models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
