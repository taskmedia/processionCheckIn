package admin

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func ExampledataHandler(c *gin.Context) {
	err := db.ResetDbExampledata()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Error loading example data",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error loading example data")

		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "Example data loaded",
	})
}
