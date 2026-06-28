package models

import "gorm.io/gorm"

type ForestWoodType struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `json:"image_url"`
	Badge       string `json:"badge"`
	SortOrder   int    `json:"sort_order"`
}
