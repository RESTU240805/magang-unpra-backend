package models

type Creed struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	TitleJP     string `json:"title_jp"`
	TitleEN     string `json:"title_en"`
	Roma        string `json:"roma"`
	Tagline     string `json:"tagline"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
}
