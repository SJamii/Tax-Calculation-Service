package routes

import (
	"github.com/SJamii/tax-calculation-service/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var route = "tax-assesment"

func TaxAssesmentRouter(router *gin.Engine) {

	router.GET(route+"/", controllers.TaxAssesmentList)
	router.POST(route+"/save", controllers.TaxAssesmentSave)
}
