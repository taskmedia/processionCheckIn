package user

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

// create user
func CreateUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error binding JSON in createUserHandler in createUserHandler")
		return
	}

	id, err := db.CreateUser(user)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error creating user in database in createUserHandler")

		return
	}

	c.IndentedJSON(201, gin.H{
		"id": id,
	})
}

// delete user
func DeleteUserHandler(c *gin.Context) {
	id, err := getIdFromParam(c, "id")
	if err != nil {
		return
	}

	err = db.DeleteUser(id)
	if err != nil {
		if err.Error() == "db.DeleteUser user does not exist" {

			c.IndentedJSON(404, gin.H{
				"id":      id,
				"message": "User not found",
			})

			log.WithFields(log.Fields{
				"id": id,
			}).Error("Error deleting not existing user in DeleteUser")

			return
		}

		c.IndentedJSON(500, gin.H{
			"message": "Internal server error",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error deleting user from database in deleteUserHandler")

		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "User deleted",
	})
}
