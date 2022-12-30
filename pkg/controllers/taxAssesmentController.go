package controllers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/SJamii/tax-calculation-service/pkg/config"
	"github.com/SJamii/tax-calculation-service/pkg/models"
	"github.com/SJamii/tax-calculation-service/pkg/repository"
	"github.com/SJamii/tax-calculation-service/pkg/response"
	"github.com/gin-gonic/gin"
)

func TaxAssesmentList(ctx *gin.Context) {
	log.Info(" ----  Into the list of books function  ---- ")
	var books []models.TaxAssesment
	result := config.DB.Find(&books)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	responseData := response.Response{
		Status:  http.StatusOK,
		Data:    books,
		Message: "Data Saved successfully",
	}

	ctx.JSON(http.StatusOK, responseData)
}

func TaxAssesmentSave(ctx *gin.Context) {

	log.Info(" ---  Tax Assesment Controller  --- ")

	body := models.TaxAssesment{}

	fmt.Println(body)

	if err := ctx.BindJSON(&body); err != nil {
		responseData := response.Response{
			Status:  http.StatusBadRequest,
			Data:    err,
			Message: "Invaild",
		}

		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusOK, responseData)
		// ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	responseData := repository.TaxAssesmentSaveRepo(ctx, body)

	// responseData := response.Response{
	// 	Status:  http.StatusOK,
	// 	Data:    body,
	// 	Message: "Data Saved successfully",
	// }
	// return responseData

	ctx.JSON(http.StatusOK, responseData)
}
