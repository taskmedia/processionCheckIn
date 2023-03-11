package backend

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api"
	"github.com/taskmedia/processionCheckIn/pkg/cmd/parameter"
)

func StartRouter() error {
	r := SetupRouter()

	return r.Run(fmt.Sprintf("%s:%d", *parameter.Host, *parameter.Port))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/version", api.VersionHandler)

	apiRouter := r.Group("/api")
	api.ApiRouters(apiRouter)

	return r
}
