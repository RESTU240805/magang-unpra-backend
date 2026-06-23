package models

import "gorm.io/gorm"

type TeamMember struct {
	gorm.Model
	Name        string `json:"name"`
	Position    string `json:"position"`
	Description string `json:"description" gorm:"type:text"`
	PhotoPath   string `json:"photo_path"`
	SortOrder   int    `json:"sort_order" gorm:"default:0"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	IsFeatured  bool   `json:"is_featured" gorm:"default:false"`
}
