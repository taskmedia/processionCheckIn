package handler

import "github.com/gin-gonic/gin"

func GroupRouters(r *gin.RouterGroup, path string, addRouter func(router *gin.RouterGroup)) {
	rg := r.Group(path)
	addRouter(rg)
}
