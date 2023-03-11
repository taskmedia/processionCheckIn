package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func ListUsersHandler(c *gin.Context) {
	generic.HandleListRequest(c, db.GetUsers)
}

func ListUserHandler(c *gin.Context) {
	generic.HandleListIdRequest(c, db.GetUser)
}
