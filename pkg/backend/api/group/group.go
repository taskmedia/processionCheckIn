package group

import "github.com/gin-gonic/gin"

func GroupsRouters(router *gin.RouterGroup) {
	router.GET("/", listGroupHandler)
	router.GET("/list", listGroupHandler)
	// router.GET("/:id", nil)

	// router.POST("/", nil)
	// router.DELETE("/:id", nil)
}
