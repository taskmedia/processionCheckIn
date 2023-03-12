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
				router.GET("/:id", NotYetImplementedHandler)

				router.POST("/", group.CreateGroupHandler)
				router.DELETE("/:id", NotYetImplementedHandler)
			},
		},
		{
			"/locations",
			func(router *gin.RouterGroup) {
				router.GET("/", location.ListLocationsHandler)
				router.GET("/list", location.ListLocationsHandler)
				router.GET("/:id", NotYetImplementedHandler)

				router.POST("/", location.CreateLocationHandler)
				router.DELETE("/:id", NotYetImplementedHandler)
			},
		},
		{
			"/seasons",
			func(router *gin.RouterGroup) {
				router.GET("/", NotYetImplementedHandler)
				router.GET("/list", NotYetImplementedHandler)
				router.GET("/:id", NotYetImplementedHandler)

				router.POST("/", NotYetImplementedHandler)
				router.DELETE("/:id", NotYetImplementedHandler)
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

func NotYetImplementedHandler(c *gin.Context) {
	c.IndentedJSON(501, gin.H{
		"message": "Not yet implemented",
	})
}
