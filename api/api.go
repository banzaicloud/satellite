package api

import (
	"github.com/banzaicloud/noaa/defaults"
	"github.com/banzaicloud/noaa/providers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Identifier is the common interface
type Identifier interface {
	Identify() (string, error)
}

// ProviderInfo holds the provider name
type ProviderInfo struct {
	Name string `json:"name"`
}

// DetermineProvider determines the cloud provider
func DetermineProvider(c *gin.Context) {

	identifiers := []Identifier{
		&providers.IdentifyAzure{Log: log},
		&providers.IdentifyAmazon{Log: log},
		&providers.IdentifyDigitalOcean{Log: log},
		&providers.IdentifyOracle{Log: log},
		&providers.IdentifyGoogle{Log: log},
		&providers.IdentifyAlibaba{Log: log},
	}
	identifiedProv := defaults.Unknown
	var err error
	for _, prov := range identifiers {
		identifiedProv, err = prov.Identify()
		if err != nil {
			log.Warn(err)
			continue
		}
		if identifiedProv != defaults.Unknown {
			c.JSON(http.StatusOK, &ProviderInfo{
				Name: identifiedProv,
			})
			return
		}
	}
	identifiedProv, _ = (&providers.IdentifySlow{Log: log}).Identify()
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
