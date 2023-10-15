package model

import "gorm.io/gorm"

type Movie struct {
	*gorm.Model
	Name           string
	Language       string
	ReleaseDate    string
	Duration       uint
	AgeRestriction uint
	Actors         string
	Director       string
	Description    string
	Completed      bool
	Liked          bool
	MyList         bool
}
