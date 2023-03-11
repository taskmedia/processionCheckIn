package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIdFromParam(c *gin.Context, param string) (int, error) {
	param_id := c.Param(param)
	id, err := strconv.Atoi(param_id)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   "Invalid id - must be an integer",
		})
		return -1, err
	}

	return id, nil
}
