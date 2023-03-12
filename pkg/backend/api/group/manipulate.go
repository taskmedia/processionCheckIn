package group

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

// create group
func CreateGroupHandler(c *gin.Context) {
	generic.HandleCreateRequests(c, db.CreateGroup)
}
