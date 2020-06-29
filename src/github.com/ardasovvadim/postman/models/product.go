package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	Name     string     `json:"name"`
	Price    float32    `json:"price"`
	Quantity int        `json:"quantity"`
	ImageUrl string     `json:"imageUrl"`
	Packages []*Package `json:"packages" gorm:"many2many:PackageProducts"`
}
