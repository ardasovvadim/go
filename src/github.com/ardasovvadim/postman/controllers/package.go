package controllers

import (
	"github.com/ardasovvadim/postman/models"
	"github.com/ardasovvadim/postman/models/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func FindPackages(c *gin.Context) {
	var packages []models.Package
	models.DB.Find(&packages)

	c.JSON(http.StatusOK, gin.H{"data": packages})
}

func FindPackage(c *gin.Context) {
	var _package models.Package

	if err := models.DB.Where("id = ?", c.Param("id")).First(&_package).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": _package})
}

func CreatePackage(c *gin.Context) {
	var model request.PackageCreateEditModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products := make([]*models.Product, len(model.Products))
	for i := range model.Products {
		var product models.Product
		models.DB.First(&product, model.Products[i])
		products[i] = &product
	}
	var _package = models.Package{
		DeliveryAddress: model.DeliveryAddress,
		Recipient:       model.Recipient,
		Sender:          model.Sender,
		DispatchDate:    time.Now(),
		Status:          "Sent",
		Products:        products,
	}
	models.DB.Create(&_package)

	c.JSON(http.StatusOK, gin.H{"data": _package})
}

func FindPackagesByRecipient(c *gin.Context) {
	var _packages []models.Package

	name := c.Param("name")
	name = "%" + strings.ReplaceAll(name, "-", " ") + "%"

	if err := models.DB.Where("recipient LIKE ?", name).Preload("Products").Find(&_packages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": _packages})
}

func SetDeliveredPackage(c *gin.Context) {
	var _package models.Package

	id := c.Param("id")

	models.DB.First(&_package, id)

	_package.Status = "Delivered"
	t := time.Now()
	_package.RecipientDate = &t

	models.DB.Save(_package)

	c.JSON(http.StatusOK, gin.H{"data": _package})
}
