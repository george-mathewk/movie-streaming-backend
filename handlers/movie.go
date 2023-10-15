package handlers

import (
	"net/http"
	"strconv"

	"github.com/george-mathewk/movie-streaming-backend/database"
	"github.com/george-mathewk/movie-streaming-backend/model"
	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context){
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"err":err,
		})
		return
	}

	resp := database.DB.Create(&movie)
	if resp.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"err":"Error in adding data to databse",
		})
		return
	}

	c.JSON(http.StatusOK, "Successfully added Movie")
}

func GetMovies(c *gin.Context){
	var movies []model.Movie

	resp := database.DB.Find(&movies)

	if resp.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Error in fetching data from the database",
		})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func GetMovie(c *gin.Context){
	var movie model.Movie

	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Error in getting movieID",
		})
		return
	}
	resp := database.DB.Where("id= ?", id).Find(&movie)
	if resp.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Movie not available",
		})
		return
	}
	if resp.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No record found for the provided id",
		})
		return
	}

	c.JSON(http.StatusOK, movie)
}