package models

type SustainabilityImage struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ImageURL string `json:"image_url"`
}
