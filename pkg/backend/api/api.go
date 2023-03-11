package api

import (
	"github.com/gin-gonic/gin"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/db"
	"github.com/taskmedia/processionCheckIn/pkg/backend/api/users"
)

func ApiRouters(router *gin.RouterGroup) {
	router.GET("/ping", pingHandler)
	router.GET("/version", VersionHandler)

	usersRouter := router.Group("/users")
	users.UsersRouters(usersRouter)

	dbRouter := router.Group("/db")
	db.DbRouters(dbRouter)
}
