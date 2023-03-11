package generic

import "github.com/gin-gonic/gin"

func HandleGroupRouters(r *gin.RouterGroup, path string, routers func(router *gin.RouterGroup)) {
	rg := r.Group(path)
	routers(rg)
}
