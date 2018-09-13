package providers

import (
	"github.com/banzaicloud/whereami/defaults"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

//Used docs
//https://www.alibabacloud.com/help/faq-detail/49122.htm

type IdentifyAlibaba struct {
	Log logrus.FieldLogger
}

func (a *IdentifyAlibaba) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return defaults.Unknown, err
	}
	if strings.Contains(string(data), "Alibaba Cloud") {
		return defaults.Alibaba, nil
	}
	return defaults.Unknown, nil
}

func IdentifyAlibabaViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	req, err := http.NewRequest("GET", "http://100.100.100.200/latest/meta-data/instance/instance-type", nil)
	if err != nil {
		log.Errorf("could not create proper http request %s", err.Error())
		detected <- defaults.Unknown
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Something happened during the request %s", err.Error())
		detected <- defaults.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("Something happened during parsing the response body %s", err.Error())
			detected <- defaults.Unknown
			return
		}
		if strings.HasPrefix(string(body), "ecs.") {
			detected <- defaults.Alibaba
		}
	}
}
