package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title         string      `json:"title"`
	Slug          string      `json:"slug" gorm:"unique"`
	Summary       string      `json:"summary"`
	Content       string      `json:"content"`
	ThumbnailPath string      `json:"thumbnail_path"`
	PublishedAt   *time.Time  `json:"published_at"`
	IsPublished   bool        `json:"is_published" gorm:"default:false"`
	Images        []NewsImage `json:"Images" gorm:"foreignKey:NewsID"`
	Category      string      `json:"category"`
}

type NewsImage struct {
	gorm.Model
	NewsID   uint   `json:"news_id"`
	ImageURL string `json:"image_url"`
}
