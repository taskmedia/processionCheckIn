package user

import "github.com/gin-gonic/gin"

func UsersRouters(router *gin.RouterGroup) {
	router.GET("/list", listHandler)
	router.GET("/:id", userHandler)
}
