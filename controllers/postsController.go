package controllers

import (
	"errors"
	"net/http"

	"github.com/Abdiooa/first-go-crud/initializers"
	"github.com/Abdiooa/first-go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func PostsCreate(c *gin.Context) {
	var requestBody CreatePostRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// post := models.Post{Title: "premier", Body: " premier body"}

	post := models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all posts"})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetAPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func UpdateAPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching post"})
		return
	}

	var input UpdatePostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Body != "" {
		post.Body = input.Body
	}

	if err := initializers.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeleteAPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching post"})
		return
	}

	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Post deleted successfully"})
}
