package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func listUsersHandler(c *gin.Context) {
	generic.HandleListRequest(c, db.GetUsers)
}

func listUserHandler(c *gin.Context) {
	id, err := getIdFromParam(c, "id")
	if err != nil {
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
