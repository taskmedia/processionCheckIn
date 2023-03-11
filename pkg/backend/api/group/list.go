package group

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func listGroupHandler(c *gin.Context) {
	groups, err := db.GetGroups()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, groups)
}
