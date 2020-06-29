package main

import (
	"github.com/ardasovvadim/postman/controllers"
	"github.com/ardasovvadim/postman/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	// packages
	r.GET("/package", controllers.FindPackages)
	r.POST("/package", controllers.CreatePackage)
	r.GET("/package/:id", controllers.FindPackage)
	r.GET("/recipient/:name/packages", controllers.FindPackagesByRecipient)
	r.PUT("/package/:id", controllers.SetDeliveredPackage)
	// products
	r.GET("/product", controllers.FindProducts)
	r.POST("/product", controllers.CreateProduct)

	r.Run()
}
