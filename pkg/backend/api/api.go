package api

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/db"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/group"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/location"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/user"
)

func ApiRouters(router *gin.RouterGroup) {
	router.GET("/ping", pingHandler)
	router.GET("/version", VersionHandler)

	dbRouter := router.Group("/db")
	db.DbRouters(dbRouter)

	groupsRouter := router.Group("/groups")
	group.GroupsRouters(groupsRouter)

	locationsRouter := router.Group("/locations")
	location.LocationsRouters(locationsRouter)

	usersRouter := router.Group("/users")
	user.UsersRouters(usersRouter)
}
