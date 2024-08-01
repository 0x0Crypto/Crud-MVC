package models

import (
	controller "mvc/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes this function init routes
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", controller.ListUsers)
	r.GET("/users/:id", controller.FindUser)
	r.POST("/users", controller.NewUser)
}
