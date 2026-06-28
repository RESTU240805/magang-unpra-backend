package models

import "gorm.io/gorm"

type SupplyChainPolicy struct {
	gorm.Model
	Title      string `json:"title"`
	Points     string `json:"points" gorm:"type:text"`
	Procedures string `json:"procedures" gorm:"type:text"`
	SortOrder  int    `json:"sort_order"`
	IsActive   bool   `json:"is_active" gorm:"default:true"`
}
