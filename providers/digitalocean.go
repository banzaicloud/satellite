package providers

import (
	"encoding/json"
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

//Used docs
// https://developers.digitalocean.com/documentation/metadata/#metadata-in-json

type digitalOceanMetadataResponse struct {
	DropletID int `json:"droplet_id"`
}

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

func IdentifyDigitalOceanViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	r := digitalOceanMetadataResponse{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/v1.json", nil)
	if err != nil {
		log.Errorf("could not create a proper http request %s", err.Error())
		detected <- api.Unknown
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Something happened during the request %s", err.Error())
		detected <- api.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("Something happened during parsing the response body %s", err.Error())
			detected <- api.Unknown
			return
		}
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Errorf("Something happened during unmarshalling the response body %s", err.Error())
			detected <- api.Unknown
			return
		}
		if r.DropletID > 0 {
			detected <- api.DigitalOcean
		}
	}
}
