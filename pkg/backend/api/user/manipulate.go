package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

// create user
func CreateUserHandler(c *gin.Context) {
	generic.HandleCreateRequests(c, db.CreateUser)
}

// delete user
func DeleteUserHandler(c *gin.Context) {
	generic.HandleDeleteIdRequest(c, db.DeleteUser)
}
