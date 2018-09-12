package providers

import (
	"encoding/json"
	"fmt"
	"github.com/banzaicloud/whereami/api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

// Used docs
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

type instanceIdentityResponse struct {
	ImageID    string `json:"imageId"`
	InstanceID string `json:"instanceId"`
}

type IdentifyAmazon struct {
	log logrus.FieldLogger
}

func (a *IdentifyAmazon) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/product_version")
	if err != nil {
		logrus.Errorf("Something happened during reading a file: %s", err.Error())
		return "", err
	}
	if strings.Contains(string(data), "amazon") {
		return api.Amazon, nil
	}
	return api.Unknown, nil
}

func IdentifyAmazonViaMetadataServer() error {
	//r := instanceIdentityResponse{}
	//req, err := http.NewRequest("GET", "http://169.254.169.254/latest/dynamic/instance-identity/document", nil)
	//if err != nil {
	//   a.log.Errorf("could not create a new amazon request %s", err.Error())
	//}
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//    a.log.Errorf("Something happened during the request %s", err.Error())
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//    a.log.Errorf("Something happened during parsing the response body %s", err.Error())
	//}
	//err = json.Unmarshal(body, &r)
	//if err != nil {
	//    a.log.Errorf("Something happened during unmarshalling the response body %s", err.Error())
	//}
}
