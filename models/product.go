package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ProductName  string `json:"product_name"`
	Type         string `json:"type"`
	ProductImage string `json:"product_image"`
	Description  string `json:"description"`
}
