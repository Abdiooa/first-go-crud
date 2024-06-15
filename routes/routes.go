package routes

import (
	"github.com/Abdiooa/first-go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetAPost)
	r.PATCH("/posts/:id", controllers.UpdateAPost)
	r.DELETE("/posts/:id", controllers.DeleteAPost)
}
