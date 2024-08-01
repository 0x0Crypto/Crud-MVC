package main

import (
	"mvc/initializers"
	models "mvc/models/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	models.InitRoutes(&router.RouterGroup)
	router.Run(":8080")
}
