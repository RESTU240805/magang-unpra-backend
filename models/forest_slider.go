package models

import "gorm.io/gorm"

type ForestSlider struct {
	gorm.Model
	ImageURL  string `json:"image_url"`
	SortOrder int    `json:"sort_order"`
}
