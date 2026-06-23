package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string         `json:"name"`
	Summary       string         `json:"summary"`
	Description   string         `json:"description"`
	ThumbnailPath string         `json:"thumbnail_path"`
	Category      string         `json:"category"`
	Tags          string         `json:"tags"`
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	Images        []ProductImage `json:"Images" gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}
