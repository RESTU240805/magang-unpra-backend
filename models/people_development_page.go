package models

import "gorm.io/gorm"

type PeopleDevelopmentPage struct {
	gorm.Model
	Description string `json:"description" gorm:"type:text"`
}
