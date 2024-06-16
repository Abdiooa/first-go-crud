package main

import (
	"github.com/Abdiooa/first-go-crud/initializers"
	"github.com/Abdiooa/first-go-crud/models"
	"github.com/Abdiooa/first-go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	r := gin.Default()
	r.Use(initializers.CorsMiddleware())
	routes.SetupRoutes(r)
	r.Run()
}
