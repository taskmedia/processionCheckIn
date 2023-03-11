package db

import "github.com/gin-gonic/gin"

func DbRouters(router *gin.RouterGroup) {
	router.GET("/init", initHandler)
	router.GET("/reset", resetHandler)
	router.GET("/exampledata", exampledataHandler)
}
