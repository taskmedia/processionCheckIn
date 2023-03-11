package user

import "github.com/gin-gonic/gin"

func UsersRouters(router *gin.RouterGroup) {
	router.GET("/", listUsersHandler)
	router.GET("/list", listUsersHandler)
	router.GET("/:id", listUserHandler)

	router.POST("/", createUserHandler)
	router.DELETE("/:id", deleteUserHandler)
}
