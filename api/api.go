package api

import (
	"github.com/banzaicloud/whereami/defaults"
	"github.com/banzaicloud/whereami/providers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Identifier interface {
	Identify() (string, error)
}

type ProviderInfo struct {
	Name string `json:"name"`
}

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
