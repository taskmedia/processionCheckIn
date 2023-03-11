package generic

import "github.com/gin-gonic/gin"

func HandleListRequest(c *gin.Context, getDataFunc func() (interface{}, error)) {
	data, err := getDataFunc()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, data)
}
