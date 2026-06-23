package models

import "time"

type Sustainability struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CoverImage  string    `json:"cover_image"`
	Icon        string    `json:"icon"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `gorm:"type:boolean" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
