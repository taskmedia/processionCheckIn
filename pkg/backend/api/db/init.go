package db

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func InitHandler(c *gin.Context) {
	c.IndentedJSON(501, gin.H{
		"message": "Not implemented",
	})

	db.InitDb()
}
