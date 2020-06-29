package controllers

import (
	"github.com/ardasovvadim/postman/models"
	"github.com/ardasovvadim/postman/models/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func FindProduct(c *gin.Context) {

}

func CreateProduct(c *gin.Context) {
	var model request.ProductCreateEditModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:     model.Name,
		Price:    model.Price,
		Quantity: model.Quantity,
		ImageUrl: model.ImageUrl,
	}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
