package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Package struct {
	gorm.Model

	DeliveryAddress string     `json:"deliveryAddress"`
	Recipient       string     `json:"recipient"`
	Sender          string     `json:"sender"`
	DispatchDate    time.Time  `json:"dispatchDate"`
	RecipientDate   *time.Time `json:"recipientDate"`
	Status          string     `json:"status" gorm:"default:'Sent'"`
	Products        []*Product `json:"products" gorm:"many2many:PackageProducts"`
}
