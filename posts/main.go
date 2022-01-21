package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salah-jr/microservices/DB"
	"net/http"
)

func main() {

	g := gin.Default()
	posts := g.Group("posts")
	{
		posts.GET("/", Welcome)
		posts.GET("/my-posts", MyPosts)

	}
	g.Run("127.0.0.1:8090")
}

func MyPosts(g *gin.Context) {
	userID := g.GetHeader("USER_ID")
	var posts []DB.Post
	DB.Con.Where("user_id", userID).Find(&posts)

	if len(posts) > 0 {
		g.JSON(http.StatusOK, gin.H{
			"message": "Here's your posts!",
			"status":  true,
			"data":    posts,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"message": "No posts found",
			"status":  true,
			"data":    "",
		})
	}
}

func Welcome(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Welcome in post microservice",
		"status":  true,
		"data":    "",
	})
}
