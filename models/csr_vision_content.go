package models

import "gorm.io/gorm"

type CsrVisionContent struct {
	gorm.Model
	CorporateDescription string `json:"corporate_description" gorm:"type:text"`
	Objectives           string `json:"objectives" gorm:"type:text"`
}
