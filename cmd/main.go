package main

import (
	"net/http"

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
	viper.SetConfigFile("./pkg/config/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	// dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World  !!!",
		})
	})
	router.Run(port)
}
