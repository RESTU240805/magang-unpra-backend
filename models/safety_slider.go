package models

import "gorm.io/gorm"

type SafetySlider struct {
	gorm.Model
	ImageURL  string `json:"image_url"`
	SortOrder int    `json:"sort_order"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}
