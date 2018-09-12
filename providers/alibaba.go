package providers

import (
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type IdentifyAlibaba struct {
	log logrus.FieldLogger
}

func (a *IdentifyAlibaba) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		logrus.Errorf("Something happened during reading a file: %s", err.Error())
		return "", err
	}
	if strings.Contains(string(data), "Alibaba Cloud") {
		return api.Alibaba, nil
	}
	return api.Unknown, nil
}
