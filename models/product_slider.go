package models

import "gorm.io/gorm"

type ProductSlider struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `json:"sort_order"`
}
