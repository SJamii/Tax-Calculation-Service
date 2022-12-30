package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaxAssesmentLineItem struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	TaxAmount         *int      `json:"taxAmount" binding:"required" validate:"required"`
	IsPercentage      bool      `json:"isPercentage" binding:"required"`
	DependentOn       *uuid.UUID
	SalaryTypeId      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	TaxAssesmentRefer uuid.UUID
}
