package entity

import (
	"gorm.io/gorm"
)

type ShippingOrder struct {
	gorm.Model

	OrderID string `json:"order_id" gorm:"index"`
	Vendor  string `json:"vendor"`
	Address string `json:"address"`
}
