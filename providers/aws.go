package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/banzaicloud/satellite/defaults"
	"github.com/sirupsen/logrus"
)

// Used docs
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

type instanceIdentityResponse struct {
	ImageID    string `json:"imageId"`
	InstanceID string `json:"instanceId"`
}

// IdentifyAmazon stuct holds the logger
type IdentifyAmazon struct {
	Log logrus.FieldLogger
}

// Identify tries to identify Amazon provider by reading the /sys/class/dmi/id/product_version file
func (a *IdentifyAmazon) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/product_version")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return defaults.Unknown, err
	}
	if strings.Contains(string(data), "amazon") {
		return defaults.Amazon, nil
	}
	return defaults.Unknown, nil
}

// IdentifyAmazonViaMetadataServer tries to identify Amazon via metadata server
func IdentifyAmazonViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	r := instanceIdentityResponse{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/latest/dynamic/instance-identity/document", nil)
	if err != nil {
		log.Errorf("could not create a new amazon request %s", err.Error())
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
		if strings.HasPrefix(r.ImageID, "ami-") &&
			strings.HasPrefix(r.InstanceID, "i-") {
			detected <- defaults.Amazon
			return
		}
	}

}
