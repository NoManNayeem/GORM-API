package routes

import (
	"GORM_API/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:code", controllers.GetProductByCode)
	r.DELETE("/product/:id", controllers.DeleteProductByID) // Add this line
}
