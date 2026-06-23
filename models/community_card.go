package models

import "gorm.io/gorm"

type CommunityCard struct {
	gorm.Model
	Title       string                 `json:"title"`
	Description string                 `json:"description" gorm:"type:text"`
	IconPath    string                 `json:"icon_path"`
	Link        string                 `json:"link" gorm:"default:'/sustainability/csr/community'"`
	SortOrder   int                    `json:"sort_order" gorm:"default:0"`
	IsActive    bool                   `json:"is_active" gorm:"default:true"`
	Images      []CommunityCardImage   `json:"Images" gorm:"foreignKey:CardID"`
}

type CommunityCardImage struct {
	gorm.Model
	CardID   uint   `json:"card_id"`
	ImageURL string `json:"image_url"`
}
