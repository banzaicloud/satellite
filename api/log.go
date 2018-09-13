package api

import (
	"github.com/banzaicloud/noaa/config"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger

func init() {
	log = config.Logger()
}
