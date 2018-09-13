package providers

import (
	"github.com/banzaicloud/whereami/defaults"
	"github.com/sirupsen/logrus"
)

// IdentifySlow struct holds the logger
type IdentifySlow struct {
	Log logrus.FieldLogger
}

// Identify tries to identify the provider via calling the metadata service
func (s *IdentifySlow) Identify() (string, error) {

	detected := make(chan string)
	defer close(detected)

	prov := []func(chan<- string, logrus.FieldLogger){
		IdentifyOracleViaMetadataServer,
		IdentifyDigitalOceanViaMetadataServer,
		IdentifyAlibabaViaMetadataServer,
		IdentifyAmazonViaMetadataServer,
		IdentifyAzureViaMetadataServer,
		IdentifyGoogleViaMetadataServer,
	}

	for _, functions := range prov {
		go functions(detected, s.Log)
	}
	for range prov {
		d := <-detected
		if d != defaults.Unknown {
			return d, nil
		}
	}
	return defaults.Unknown, nil
}
