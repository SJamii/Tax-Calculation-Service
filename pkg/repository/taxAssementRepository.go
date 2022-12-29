package repository

import (
	"net/http"

	"github.com/SJamii/tax-calculation-service/pkg/config"
	"github.com/SJamii/tax-calculation-service/pkg/models"
	"github.com/gin-gonic/gin"
)

func TaxAssesmentSaveRepo(ctx *gin.Context, data models.TaxAssesment) {

	result := config.DB.Create(&data)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}
