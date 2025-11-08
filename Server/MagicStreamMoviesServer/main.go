package main

import (
	"fmt"

	controllers "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	router.GET("/movies", controllers.GetMovies())

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
