package routes

import (
	"MyGram3/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("users/:id", controllers.DeleteUser)

	r.GET("/photos", controllers.FindPhotos)
	r.POST("/photos", controllers.CreatePhoto)
	r.GET("/photos/:id", controllers.FindPhoto)
	r.PATCH("/photos/:id", controllers.UpdatePhoto)
	r.DELETE("photos/:id", controllers.DeletePhoto)

	r.GET("/comments", controllers.FindComs)
	r.POST("/comments", controllers.CreateCom)
	r.GET("/comments/:id", controllers.FindCom)
	r.PATCH("/comments/:id", controllers.UpdateCom)
	r.DELETE("comments/:id", controllers.DeleteCom)

	r.GET("/sosmed", controllers.FindSoss)
	r.POST("/sosmed", controllers.CreateSos)
	r.GET("/sosmed/:id", controllers.FindSos)
	r.PATCH("/sosmed/:id", controllers.UpdateSos)
	r.DELETE("sosmed/:id", controllers.DeleteSos)

	return r
}
