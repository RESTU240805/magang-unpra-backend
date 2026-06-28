package models

import "gorm.io/gorm"

type ForestApproach struct {
	gorm.Model
	Description string `json:"description" gorm:"type:text"`
}
