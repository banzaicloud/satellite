package providers

import (
	"github.com/banzaicloud/whereami/defaults"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

//Used docs
// https://azure.microsoft.com/en-us/blog/announcing-general-availability-of-azure-instance-metadata-service/

// IdentifyAzure struct holds the logger
type IdentifyAzure struct {
	Log logrus.FieldLogger
}

// Identify tries to identify Azure provider by reading the /sys/class/dmi/id/sys_vendor file
func (a *IdentifyAzure) Identify() (string, error) {
	data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		a.Log.Errorf("Something happened during reading a file: %s", err.Error())
		return defaults.Unknown, err
	}
	if strings.Contains(string(data), "Microsoft Corporation") {
		return defaults.Azure, nil
	}
	return defaults.Unknown, nil
}

// IdentifyAzureViaMetadataServer tries to identify Azure via metadata server
func IdentifyAzureViaMetadataServer(detected chan<- string, log logrus.FieldLogger) {
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/instance?api-version=2017-12-01", nil)
	if err != nil {
		log.Errorf("Could not create a proper http request %s", err.Error())
		detected <- defaults.Unknown
		return
	}
	req.Header.Set("Metadata", "true")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Something happened during the request %s", err.Error())
		detected <- defaults.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		detected <- defaults.Azure
	}
}
