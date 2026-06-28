package models

import "gorm.io/gorm"

type CsrReport struct {
	gorm.Model
	Year      int    `json:"year"`
	Quarter   string `json:"quarter"`
	Period    string `json:"period"`
	FileURL   string `json:"file_url"`
	SortOrder int    `json:"sort_order"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}
