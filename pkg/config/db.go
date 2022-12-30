package config

import (
	"log"

	"github.com/SJamii/tax-calculation-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.TaxAssesment{}, &models.TaxAssesmentLineItem{})

	DB = db
}
