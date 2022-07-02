package main

import (
	"Sirka/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, "Hello World! I'am Gin Gonic")
	})

	r.POST("/DisplayUser", func(c *gin.Context) {

		result := controllers.DisplayUser(c.Request)
		c.JSON(200, result)
	})
	r.POST("/DisplayAllUsers", func(c *gin.Context) {
		result := controllers.DisplayAllUsers(c.Request)
		c.JSON(200, result)
	})

	r.Run("localhost:9090")
}
