package controllers

import (
	"MyGram3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateSosInput struct {
	Name             string `json:"name"`
	Sosial_media_url string `json:"Sosial_media_url"`
	User_id          int64  `json:"userid"`
}

type UpdateSosInput struct {
	Name             string `json:"name"`
	Sosial_media_url string `json:"Sosial_media_url"`
	User_id          int64  `json:"userid"`
}

// GET /Soss
// Get all Soss
func FindSoss(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Soss []models.SosialMedia
	db.Find(&Soss)

	c.JSON(http.StatusOK, gin.H{"data": Soss})
}

// POST /Soss
// Create new Sos
func CreateSos(c *gin.Context) {
	// Validate input
	var input CreateSosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Sos
	Sos := models.SosialMedia{Name: input.Name, Sosial_media_url: input.Sosial_media_url, User_id: input.User_id}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Sos)

	c.JSON(http.StatusOK, gin.H{"data": Sos})
}

// GET /Soss/:id
// Find a Sos
func FindSos(c *gin.Context) { // Get model if exist
	var Sos models.SosialMedia

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Sos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Sos})
}

// PATCH /Soss/:id
// Update a Sos
func UpdateSos(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Sos models.SosialMedia
	if err := db.Where("id = ?", c.Param("id")).First(&Sos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateSosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.SosialMedia
	updatedInput.Name = input.Name
	updatedInput.Sosial_media_url = input.Sosial_media_url
	updatedInput.User_id = input.User_id

	db.Model(&Sos).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Sos})
}

// DELETE /Soss/:id
// Delete a Sos
func DeleteSos(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var sosial models.SosialMedia
	if err := db.Where("id = ?", c.Param("id")).First(&sosial).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&sosial)

	c.JSON(http.StatusOK, gin.H{"data": "data sudah di hapus"})
}
