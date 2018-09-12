package providers

import (
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type IdentifyOracle struct {
	Log logrus.FieldLogger
}

func (a *IdentifyOracle) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/chassis_asset_tag")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return api.Unknown, err
	}
	if strings.Contains(string(data), "OracleCloud") {
		return api.Oracle, nil
	}
	return api.Unknown, nil
}
