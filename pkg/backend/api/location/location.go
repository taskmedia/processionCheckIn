package location

import "github.com/gin-gonic/gin"

func LocationsRouters(router *gin.RouterGroup) {
	router.GET("/", listLocationHandler)
	router.GET("/list", listLocationHandler)
	// router.GET("/:id", nil)

	// router.POST("/", nil)
	// router.DELETE("/:id", nil)
}
