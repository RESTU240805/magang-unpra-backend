package models

import "gorm.io/gorm"

type AboutSection struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	ImagePath   string `json:"image_path"`
	BadgeNumber string `json:"badge_number"`
	BadgeLabel  string `json:"badge_label"`
}
