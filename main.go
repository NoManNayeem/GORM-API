package main

import (
	"GORM_API/configs"
	"GORM_API/models"
	"GORM_API/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.SetupDatabase()
	router := gin.Default()
	routes.SetupRoutes(router)

	// AutoMigrate the database schema
	configs.DB.AutoMigrate(&models.Product{}) // Pass your model structs here

	router.Run("localhost:8083")
}
