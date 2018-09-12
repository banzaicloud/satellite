package api

import (
	"github.com/gin-gonic/gin"
)

type Identifier interface {
	Identify() (string, error)
}

const (
	Amazon       = "amazon"
	Alibaba      = "alibaba"
	Azure        = "azure"
	Google       = "google"
	Oracle       = "oracle"
	DigitalOcean = "digitalocean"
	Unknown      = "unknown"
)

func DetermineProvider(c *gin.Context) {

}
