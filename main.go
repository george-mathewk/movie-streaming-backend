package main

import (
	"net/http"

	"github.com/george-mathewk/movie-streaming-backend/database"
	"github.com/george-mathewk/movie-streaming-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.DBInit()
	if err != nil{
		panic(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/movie", handlers.GetMovies)
	r.GET("/movie/:id", handlers.GetMovie)
	r.POST("/movie", handlers.AddMovie)

	r.Run(":8080")
}
