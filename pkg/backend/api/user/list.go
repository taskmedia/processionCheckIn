package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func listHandler(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, users)
}

func userHandler(c *gin.Context) {
	param_id := c.Param("id")
	id, err := strconv.Atoi(param_id)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   "Invalid id - must be an integer",
		})
		return
	}

	user, err := db.GetUser(id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.IndentedJSON(404, gin.H{
				"message": "Not found",
				"error":   "User not found",
				"id":      id,
			})
			return
		}

		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
		})

		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error getting user from database in userHandler")

		return
	}

	c.IndentedJSON(200, user)
}
