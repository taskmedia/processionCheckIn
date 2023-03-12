package generic

import "github.com/gin-gonic/gin"

func HandlerGroupRouters(r *gin.RouterGroup, path string, addRouter func(router *gin.RouterGroup)) {
	rg := r.Group(path)
	addRouter(rg)
}
