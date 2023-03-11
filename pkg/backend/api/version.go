package api

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/version"
)

func VersionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":  version.VERSION,
		"revision": version.REVISION,
		"commit":   version.COMMIT,
	})
}
