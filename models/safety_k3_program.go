package models

import "gorm.io/gorm"

type SafetyK3Program struct {
	gorm.Model
	Description string `json:"description" gorm:"type:text"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}
