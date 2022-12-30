package repository

import (
	"fmt"
	"net/http"

	"github.com/SJamii/tax-calculation-service/pkg/config"
	"github.com/SJamii/tax-calculation-service/pkg/models"
	"github.com/SJamii/tax-calculation-service/pkg/response"
	"github.com/gin-gonic/gin"
)

func TaxAssesmentSaveRepo(ctx *gin.Context, data models.TaxAssesment) any {

	result := config.DB.Create(&data)
	fmt.Println("r", data.TaxAssesmentLineItems[0].DeletedAt)
	if result.Error != nil {
		fmt.Println("dhuckhi")
		// ctx.AbortWithError(http.StatusNotFound, result.Error)
		responseData := response.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "Invaild",
		}
		return responseData
	}
	responseData := response.Response{
		Status:  http.StatusOK,
		Data:    data,
		Message: "Data Saved successfully",
	}
	return responseData
}
