package main

import (
	"MyGram3/models"
	"MyGram3/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SosialMedia{})

	r := routes.SetupRoutes(db)
	r.Run()
}
