package main

import (
	"fmt"

	routesHandler "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	routesHandler.SetupPublicRoutes(router)
	routesHandler.SetupProtectedRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
