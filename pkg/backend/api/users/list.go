package users

import "github.com/gin-gonic/gin"

func listHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "Not implemented",
	})
}
