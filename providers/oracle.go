package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/banzaicloud/satellite/defaults"
	"github.com/sirupsen/logrus"
)

//Used doc
//https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/gettingmetadata.htm

type oracleMetadataResponse struct {
	OkeTM string `json:"oke-tm"`
}

// IdentifyOracle struct holds the logger
type IdentifyOracle struct {
	Log logrus.FieldLogger
}

// Identify tries to identify Oracle provider by reading the /sys/class/dmi/id/chassis_asset_tag file
func (a *IdentifyOracle) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/chassis_asset_tag")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return defaults.Unknown, err
	}
	if strings.Contains(string(data), "OracleCloud") {
		return defaults.Oracle, nil
	}
	return defaults.Unknown, nil
}

// IdentifyOracleViaMetadataServer tries to identify Oracle via metadata server
func IdentifyOracleViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	r := oracleMetadataResponse{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/opc/v1/instance/metadata/", nil)
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
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Errorf("Something happened during unmarshalling the response body %s", err.Error())
			detected <- defaults.Unknown
			return
		}
		if strings.Contains(r.OkeTM, "oke") {
			detected <- defaults.Oracle
		}
	} else {
		log.Errorf("Something happened during the request with status %s", resp.Status)
		detected <- defaults.Unknown
		return
	}
}
