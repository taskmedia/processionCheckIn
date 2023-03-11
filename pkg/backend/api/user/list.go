package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func listHandler(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, users)
}
