package main

import (
	"github.com/banzaicloud/noaa/api"
	"github.com/banzaicloud/noaa/config"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := config.Logger()
	logger.Info("Noaa initialization")

	router := gin.Default()

	a := api.NewDetermineProviderApi(logger)

	router.GET("/noaa", a.DetermineProvider)
	router.Run(":8888")
}
