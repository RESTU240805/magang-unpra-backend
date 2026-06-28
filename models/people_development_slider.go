package models

import "gorm.io/gorm"

type PeopleDevelopmentSlider struct {
	gorm.Model
	ImageURL  string `json:"image_url"`
	Caption   string `json:"caption"`
	SortOrder int    `json:"sort_order"`
}
