package routes

import (
	controllers "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/controller"
	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.Engine) {
	router.GET("/movies", controllers.GetMovies())
	router.POST("/registerUser", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
}
