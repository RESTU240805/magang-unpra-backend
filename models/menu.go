package models

import "time"

type Menu struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	ParentID  *uint     `json:"parent_id"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
