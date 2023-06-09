package admin

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func InitHandler(c *gin.Context) {
	err := db.InitDb()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Error initializing database",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error initializing database")

		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "Database initialized",
	})
}
