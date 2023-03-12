package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerListIdRequest(c *gin.Context, listIdFunc func(id int) (interface{}, error)) {
	id, err := getIdFromParam(c, "id")
	if err != nil {
		return
	}

	data, err := listIdFunc(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.IndentedJSON(404, gin.H{
				"message": "Not found",
				"error":   "Requested resource not found",
				"id":      id,
			})
			return
		}

		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   "Error handling list id request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error getting list item from database in HandleListIdRequest")

		return
	}

	c.IndentedJSON(200, data)
}

func HandlerListRequest(c *gin.Context, listFunc func() (interface{}, error)) {
	data, err := listFunc()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   "Error handling list request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error handling list request")
		return
	}

	c.IndentedJSON(200, data)
}
