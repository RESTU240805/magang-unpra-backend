package models

import (
	"time"
)

type ContactInfo struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	WorkingHours string    `json:"working_hours"`
	Copyright    string    `json:"copyright"`
	HeroImage    string    `json:"hero_image"`
	MapLink      string    `json:"map_link"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
