package main

import (
	"github.com/banzaicloud/satellite/api"
	"github.com/banzaicloud/satellite/config"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := config.Logger()
	logger.Info("Satellite initialization")

	router := gin.Default()

	a := api.NewDetermineProviderApi(logger)

	router.GET("/satellite", a.DetermineProvider)
	router.Run(":8888")
}
