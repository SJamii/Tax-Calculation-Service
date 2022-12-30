package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaxAssesment struct {
	gorm.Model
	ID                    uuid.UUID              `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name                  string                 `json:"name" binding:"required" gorm:"not null"`
	Year                  int                    `json:"year" binding:"required" gorm:"not null"`
	TaxAssesmentLineItems []TaxAssesmentLineItem `gorm:"foreignKey:TaxAssesmentRefer"`
}
