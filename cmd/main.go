package main

import (
	"github.com/SJamii/tax-calculation-service/pkg/config"
	"github.com/SJamii/tax-calculation-service/pkg/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	log.Info("  ---------------       API Started !!!      ------------")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello World  !!!",
	// 	})
	// })
	// router.Run(port)
	router := gin.New()
	config.Init(dbUrl)
	routes.TaxAssesmentRouter(router)

	router.Run(port)
}
