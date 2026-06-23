package models

import "gorm.io/gorm"

type OrgChart struct {
	gorm.Model
	ImagePath string `json:"image_path"`
}
