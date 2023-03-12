package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"

	log "github.com/sirupsen/logrus"
)

func ResetHandler(c *gin.Context) {
	err := db.ResetDb()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Error resetting database",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error resetting database")

		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "Database reset",
	})
}
