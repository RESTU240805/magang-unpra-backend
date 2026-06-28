package models

import "gorm.io/gorm"

type PulpProcessRecovery struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `json:"image_url"`
	Specs       string `json:"specs" gorm:"type:text"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}
