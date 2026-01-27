package routes

import (
	controllers "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/controller"
	"github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
}
