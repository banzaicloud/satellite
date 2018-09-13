package providers

import (
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

//Used Doc
//https://cloud.google.com/compute/docs/storing-retrieving-metadata#endpoints

type IdentifyGoogle struct {
	Log logrus.FieldLogger
}

func (a *IdentifyGoogle) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return api.Unknown, err
	}
	if strings.Contains(string(data), "Google") {
		return api.Google, nil
	}
	return api.Unknown, nil
}

func IdentifyGoogleViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	req, err := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/tags", nil)
	if err != nil {
		log.Errorf("Could not create a proper http request %s", err.Error())
		detected <- api.Unknown
		return
	}
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Something happened during the request %s", err.Error())
		detected <- api.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		detected <- api.Google
	}
}
