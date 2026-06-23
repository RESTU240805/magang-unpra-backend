package models

import "time"

type CompanyDocument struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Category    string    `json:"category"`  // contoh: "Document Preview"
	DocDate     string    `json:"doc_date"`  // contoh: "April 2023"
	FileType    string    `json:"file_type"` // contoh: "PDF"
	FileSize    string    `json:"file_size"` // contoh: "1.2 MB"
	Description string    `json:"description" gorm:"type:text"`
	FileURL     string    `json:"file_url"` // path file PDF
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
