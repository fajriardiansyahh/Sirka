package main

import (
	"Sirka/controllers"
	"log"
	"net/http"
)

func main() {

	// r := gin.Default()
	// r.GET("/welcome", func(c *gin.Context) {
	// 	c.JSON(200, "Hello World! I'am Gin Gonic")
	// })
	// r.POST("/DisplayUser", func(c *gin.Context) {
	// 	c.JSON(200, controllers.DisplayUser)
	// })
	// r.POST("/DisplayAllUsers", func(c *gin.Context) {
	// 	c.JSON(200, controllers.DisplayAllUsers)
	// })
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })
	// r.Run("localhost:9090")

	mux := http.NewServeMux()
	mux.HandleFunc("/DisplayUser", controllers.DisplayUser)
	mux.HandleFunc("/DisplayAllUsers", controllers.DisplayAllUsers)
	log.Println("Starting Web Server at port: 9090")

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
