package routes

import (
	controllers "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupPublicRoutes(router *gin.Engine, client *mongo.Client) {
	router.GET("/movies", controllers.GetMovies(client))
	router.POST("/registerUser", controllers.RegisterUser(client))
	router.POST("/login", controllers.LoginUser(client))
	router.POST("/logout", controllers.Logout(client))
	router.POST("/refresh", controllers.RefreshToken(client))
}
