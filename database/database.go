package database

import (
	"github.com/george-mathewk/movie-streaming-backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() error {
	db, err := gorm.Open(sqlite.Open("movie.db"), &gorm.Config{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Movie{})
	if err != nil{
		return err
	}
	DB = db
	return nil
}
