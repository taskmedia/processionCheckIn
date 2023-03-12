package location

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

// create location
func CreateLocationHandler(c *gin.Context) {
	generic.HandleCreateRequests(c, db.CreateLocation)
}
