package models

import "gorm.io/gorm"

type ProductPage struct {
	gorm.Model
	Description string `json:"description" gorm:"type:text"`
}
