package models

import "gorm.io/gorm"

type PeopleDevelopmentPillar struct {
	gorm.Model
	Title     string `json:"title"`
	Desc      string `json:"desc" gorm:"type:text"`
	SortOrder int    `json:"sort_order"`
}
