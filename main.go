package main

import (
	"fmt"
	"log"
	"mvc/initializers"
	models "mvc/models/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := gin.Default()
	models.InitRoutes(&router.RouterGroup)
	router.Run(fmt.Sprintf(":%s", port))
}
