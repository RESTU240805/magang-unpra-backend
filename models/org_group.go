package models

import "gorm.io/gorm"

type OrgGroup struct {
	gorm.Model
	Label     string    `json:"label"`
	Color     string    `json:"color" gorm:"default:blue"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	Nodes     []OrgNode `json:"nodes,omitempty" gorm:"foreignKey:GroupID"`
}
