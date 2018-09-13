package main

import (
	"github.com/banzaicloud/noaa/api"
	"github.com/banzaicloud/noaa/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Common logger for package
var log *logrus.Logger
var logger *logrus.Entry

func initLog() *logrus.Entry {
	log = config.Logger()
	logger := log.WithFields(logrus.Fields{"state": "init"})
	return logger
}

func main() {

	logger = initLog()
	logger.Info("WhereAmI initialization")
	router := gin.Default()

	router.GET("/noaa", api.DetermineProvider)
	router.Run(":8888")
}
