package providers

import (
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type IdentifyDigitalOcean struct {
	Log logrus.FieldLogger
}

func (a *IdentifyDigitalOcean) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return api.Unknown, err
	}
	if strings.Contains(string(data), "DigitalOcean") {
		return api.DigitalOcean, nil
	}
	return api.Unknown, nil
}
