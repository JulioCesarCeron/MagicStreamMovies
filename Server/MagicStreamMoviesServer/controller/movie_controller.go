package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	database "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/database"
	models "github.com/JulioCesarCeron/MagicStreamMovies/Server/MagicStreamMoviesServer/Server/MagicStreamMoviesServer/models"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []models.Movie
		
		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching movies from database"})
		}

		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			fmt.Printf("Error decoding movies: %v\n", err)

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding movies"})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}
