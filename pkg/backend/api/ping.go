package api

import (
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}
