package generic

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleListRequest(c *gin.Context, getDataFunc func() (interface{}, error)) {
	data, err := getDataFunc()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   "Error handling list request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error handling list request")
		return
	}

	c.IndentedJSON(200, data)
}
