package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Uuid        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name" gorm:"not null" binding:"required"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"not null" binding:"required,min=0"`
	Stock       int       `json:"stock" gorm:"not null;default:0" binding:"min=0"`
	Category    string    `json:"category" gorm:"not null" binding:"required"`
	Brand       string    `json:"brand"`
	SKU         string    `json:"sku" gorm:"unique"`
	ImagePath   string    `json:"image_path"`
	ImageURL    string    `json:"image_url"`
	Status      string    `json:"status" gorm:"default:active"` // active, inactive, discontinued
	CreatedBy   uuid.UUID `json:"created_by" gorm:"type:uuid"`
	UpdatedBy   uuid.UUID `json:"updated_by" gorm:"type:uuid"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	Brand       string    `json:"brand"`
	SKU         string    `json:"sku"`
	ImageURL    string    `json:"image_url"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	Category    string  `json:"category" binding:"required"`
	Brand       string  `json:"brand"`
	SKU         string  `json:"sku"`
	Status      string  `json:"status"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int    `json:"stock,omitempty"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	SKU         string  `json:"sku"`
	Status      string  `json:"status"`
}