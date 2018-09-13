package providers

import (
	"encoding/json"
	"github.com/banzaicloud/whereami/defaults"
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

// IdentifyDigitalOcean struct holds the logger
type IdentifyDigitalOcean struct {
	Log logrus.FieldLogger
}

// Identify tries to identify DigitalOcean provider by reading the /sys/class/dmi/id/sys_vendor file
func (a *IdentifyDigitalOcean) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return defaults.Unknown, err
	}
	if strings.Contains(string(data), "DigitalOcean") {
		return defaults.DigitalOcean, nil
	}
	return defaults.Unknown, nil
}

// IdentifyDigitalOceanViaMetadataServer tries to identify DigitalOcean via metadata server
func IdentifyDigitalOceanViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	r := digitalOceanMetadataResponse{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/v1.json", nil)
	if err != nil {
		log.Errorf("could not create a proper http request %s", err.Error())
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
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Errorf("Something happened during unmarshalling the response body %s", err.Error())
			detected <- defaults.Unknown
			return
		}
		if r.DropletID > 0 {
			detected <- defaults.DigitalOcean
		}
	}
}
