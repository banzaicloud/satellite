package api

import (
    "github.com/banzaicloud/whereami/config"
    "github.com/sirupsen/logrus"
)

var log logrus.FieldLogger

func init() {
    log = config.Logger()
}