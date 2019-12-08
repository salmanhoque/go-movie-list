package movie

import (
	"regexp"
	"sort"
	"strconv"
)

// Persistence is used to save and read from file or database
type Persistence interface {
	Save(list interface{}) error
	Read(list interface{}) error
}

// Repo is used to save and query movies
type Repo struct {
	MovieList []Schema
	Storage   Persistence
}

// Add movie to the list
func (r *Repo) Add(name string, year string, rating float64) (Schema, error) {
	var err error
	var newMoview Schema

	newMoview = Schema{MovieName: name, ReleaseYear: year, MovieRating: rating}
	r.MovieList = append(r.MovieList, newMoview)
	err = r.Storage.Save(r.MovieList)
	if err != nil {
		return newMoview, err
	}

	return newMoview, err
}

// SortByRating - sort movie by rating
func (r *Repo) SortByRating() {
	movies := r.MovieList

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].MovieRating > movies[j].MovieRating
	})

	r.MovieList = movies
}

// FindByYear - find movies of a year
func (r *Repo) FindByYear(year int) []Schema {
	var movies []Schema

	for _, movie := range r.MovieList {
		if releaseYear, _ := strconv.Atoi(movie.ReleaseYear); releaseYear == year {
			movies = append(movies, movie)
		}
	}

	return movies
}

// FindByTitle - find movie with a keyword search
func (r *Repo) FindByTitle(name string) []Schema {
	var movies []Schema
	re := regexp.MustCompile("(?i)" + name)

	for _, movie := range r.MovieList {
		if re.Match([]byte(movie.MovieName)) {
			movies = append(movies, movie)
		}
	}

	return movies
}
