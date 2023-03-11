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
			func(router *gin.RouterGroup) {
				router.GET("/init", db.InitHandler)
				router.GET("/reset", db.ResetHandler)
				router.GET("/exampledata", db.ExampledataHandler)
			},
		},
		{
			"/groups",
			func(router *gin.RouterGroup) {
				router.GET("/", group.ListGroupsHandler)
				router.GET("/list", group.ListGroupsHandler)
				// router.GET("/:id", nil)

				// router.POST("/", nil)
				// router.DELETE("/:id", nil)
			},
		},
		{
			"/locations",
			func(router *gin.RouterGroup) {
				router.GET("/", location.ListLocationsHandler)
				router.GET("/list", location.ListLocationsHandler)
				// router.GET("/:id", nil)

				// router.POST("/", nil)
				// router.DELETE("/:id", nil)}
			},
		},
		{
			"/users",
			func(router *gin.RouterGroup) {
				router.GET("/", user.ListUsersHandler)
				router.GET("/list", user.ListUsersHandler)
				router.GET("/:id", user.ListUserHandler)

				router.POST("/", user.CreateUserHandler)
				router.DELETE("/:id", user.DeleteUserHandler)
			},
		},
	}

	for _, rg := range routerGroups {
		generic.HandleGroupRouters(router, rg.path, rg.routers)
	}
}
