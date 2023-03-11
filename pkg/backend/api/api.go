package api

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/db"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/group"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/location"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/user"
)

func ApiRouters(router *gin.RouterGroup) {
	router.GET("/ping", pingHandler)
	router.GET("/version", VersionHandler)

	var routerGroups = []struct {
		path    string
		routers func(router *gin.RouterGroup)
	}{
		{
			"/db",
			db.DbRouters,
		},
		{
			"/groups",
			group.GroupsRouters,
		},
		{
			"/locations",
			location.LocationsRouters,
		},
		{
			"/users",
			user.UsersRouters,
		},
	}

	for _, rg := range routerGroups {
		generic.HandleGroupRouters(router, rg.path, rg.routers)
	}
}
