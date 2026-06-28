package models

import "gorm.io/gorm"

type SafetyPolicy struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Icon        string `json:"icon"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}
