package generic

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleCreateRequests(c *gin.Context, createFunc func(*gin.Context) (int, error)) {
	id, err := createFunc(c)
	if err != nil {
		return
	}

	c.IndentedJSON(201, gin.H{
		"id": id,
	})
}

func HandleDeleteIdRequest(c *gin.Context, deleteFunc func(id int) error) {
	id, err := getIdFromParam(c, "id")
	if err != nil {
		return
	}

	err = deleteFunc(id)
	if err != nil {
		if err.Error() == "id does not exist" {
			c.IndentedJSON(404, gin.H{
				"message": "Not found",
				"error":   "Requested resource not found",
				"id":      id,
			})

			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Error deleting item from database in HandleDeleteIdRequest")

			return
		}

		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   "Error handling delete id request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error deleting item from database in HandleDeleteIdRequest")

		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "resource deleted",
	})
}
