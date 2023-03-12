package api

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/admin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/generic"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
)

func ApiRouters(router *gin.RouterGroup) {
	router.GET("/ping", pingHandler)
	router.GET("/version", VersionHandler)

	var routerGroups = []struct {
		path    string
		routers func(router *gin.RouterGroup)
	}{
		{
			"/admin",
			func(router *gin.RouterGroup) {
				router.GET("/init", admin.InitHandler)
				router.GET("/reset", admin.ResetHandler)
				router.GET("/exampledata", admin.ExampledataHandler)
			},
		},
		{
			"/groups",
			func(router *gin.RouterGroup) {
				router.GET("/", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetGroups) })
				router.GET("/list", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetGroups) })
				router.GET("/:id", NotYetImplementedHandler)

				router.POST("/", func(c *gin.Context) { generic.HandlerCreateRequests(c, db.CreateGroup) })
				router.DELETE("/:id", NotYetImplementedHandler)
			},
		},
		{
			"/locations",
			func(router *gin.RouterGroup) {
				router.GET("/", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetLocations) })
				router.GET("/list", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetLocations) })
				router.GET("/:id", NotYetImplementedHandler)

				router.POST("/", func(c *gin.Context) { generic.HandlerCreateRequests(c, db.CreateLocation) })
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
				router.GET("/", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetUsers) })
				router.GET("/list", func(c *gin.Context) { generic.HandlerListRequest(c, db.GetUsers) })
				router.GET("/:id", func(c *gin.Context) { generic.HandlerListIdRequest(c, db.GetUser) })

				router.POST("/", func(c *gin.Context) { generic.HandlerCreateRequests(c, db.CreateUser) })
				router.DELETE("/:id", func(c *gin.Context) { generic.HandlerDeleteIdRequest(c, db.DeleteUser) })
			},
		},
	}

	for _, rg := range routerGroups {
		generic.HandlerGroupRouters(router, rg.path, rg.routers)
	}
}

func NotYetImplementedHandler(c *gin.Context) {
	c.IndentedJSON(501, gin.H{
		"message": "Not yet implemented",
	})
}
