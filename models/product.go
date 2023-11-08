package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID           string `gorm:"type:VARCHAR(50);primary key" json:"id"`
	ProductName  string `json:"product_name"`
	Type         string `json:"type"`
	ProductImage string `json:"product_image"`
	Description  string `json:"description"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now().Local().UTC()
	p.UpdatedAt = time.Now().Local().UTC()

	return
}
