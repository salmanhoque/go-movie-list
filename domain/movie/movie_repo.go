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
func (m *Repo) Add(name string, year string, rating float64) (Schema, error) {
	var err error
	var newMoview Schema

	newMoview = Schema{MovieName: name, ReleaseYear: year, MovieRating: rating}
	m.MovieList = append(m.MovieList, newMoview)
	err = m.Storage.Save(m.MovieList)
	if err != nil {
		return newMoview, err
	}

	return newMoview, err
}

// SortByRating - sort movie by rating
func (m *Repo) SortByRating() {
	movies := m.MovieList

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].MovieRating > movies[j].MovieRating
	})

	m.MovieList = movies
}

// FindByYear - find movies of a year
func (m *Repo) FindByYear(year int) []Schema {
	var movies []Schema

	for _, movie := range m.MovieList {
		if releaseYear, _ := strconv.Atoi(movie.ReleaseYear); releaseYear == year {
			movies = append(movies, movie)
		}
	}

	return movies
}

// FindByTitle - find movie with a keyword search
func (m *Repo) FindByTitle(name string) []Schema {
	var movies []Schema
	re := regexp.MustCompile("(?i)" + name)

	for _, movie := range m.MovieList {
		if re.Match([]byte(movie.MovieName)) {
			movies = append(movies, movie)
		}
	}

	return movies
}
