package location

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func listLocationHandler(c *gin.Context) {
	generic.HandleListRequest(c, db.GetLocations)
}
