package providers

import (
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type IdentifyAzure struct {
	Log logrus.FieldLogger
}

func (a *IdentifyAzure) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return api.Unknown, err
	}
	if strings.Contains(string(data), "Microsoft Corporation") {
		return api.Azure, nil
	}
	return api.Unknown, nil
}
