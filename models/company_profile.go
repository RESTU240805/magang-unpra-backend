package models

import "gorm.io/gorm"

type CompanyProfile struct {
	gorm.Model
	Title        string `json:"title"`
	Content      string `json:"content"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	HeroImage    string `json:"hero_image"`
	CreedBgImage string `json:"creed_bg_image"`
}
