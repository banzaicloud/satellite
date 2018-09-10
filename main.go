package main

import (
    "github.com/banzaicloud/whereami/api"
    "github.com/banzaicloud/whereami/config"
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

    router.GET("/whereami",api.DetermineProvider)
    router.Run()
}
