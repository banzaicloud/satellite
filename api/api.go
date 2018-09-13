package api

import (
	"net/http"

	"github.com/banzaicloud/noaa/defaults"
	"github.com/banzaicloud/noaa/providers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Identifier is the common interface
type Identifier interface {
	Identify() (string, error)
}

// ProviderInfo holds the provider name
type ProviderInfo struct {
	Name string `json:"name"`
}

// DetermineProviderApi determines the cloud provider.
type DetermineProviderApi struct {
	logger logrus.FieldLogger
}

// NewDetermineProviderApi returns a new DetermineProviderApi instance.
func NewDetermineProviderApi(logger logrus.FieldLogger) *DetermineProviderApi {
	return &DetermineProviderApi{
		logger: logger,
	}
}

// DetermineProvider determines the cloud provider
func (a *DetermineProviderApi) DetermineProvider(c *gin.Context) {

	identifiers := []Identifier{
		&providers.IdentifyAzure{Log: a.logger},
		&providers.IdentifyAmazon{Log: a.logger},
		&providers.IdentifyDigitalOcean{Log: a.logger},
		&providers.IdentifyOracle{Log: a.logger},
		&providers.IdentifyGoogle{Log: a.logger},
		&providers.IdentifyAlibaba{Log: a.logger},
	}
	identifiedProv := defaults.Unknown
	var err error
	for _, prov := range identifiers {
		identifiedProv, err = prov.Identify()
		if err != nil {
			a.logger.Warn(err)
			continue
		}
		if identifiedProv != defaults.Unknown {
			c.JSON(http.StatusOK, &ProviderInfo{
				Name: identifiedProv,
			})
			return
		}
	}
	identifiedProv, _ = (&providers.IdentifySlow{Log: a.logger}).Identify()
	if identifiedProv != defaults.Unknown {
		c.JSON(http.StatusOK, &ProviderInfo{
			Name: identifiedProv,
		})
		return
	}
	c.JSON(http.StatusNotFound, &ProviderInfo{
		Name: identifiedProv,
	})
}
