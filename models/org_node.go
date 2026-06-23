package models

import "gorm.io/gorm"

type OrgNode struct {
	gorm.Model
	GroupID   uint      `json:"group_id"`
	ParentID  *uint     `json:"parent_id" gorm:"default:null"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	PhotoPath string    `json:"photo_path"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	Children  []OrgNode `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}
