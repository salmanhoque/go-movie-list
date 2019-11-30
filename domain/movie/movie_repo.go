package domain

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

// MovieRepo is used to save and query movies
type MovieRepo struct {
	MovieList []Movie
	Storage   Persistence
}

// Add movie to the list
func (m *MovieRepo) Add(name string, year string, rating float64) (Movie, error) {
	var err error
	var newMoview Movie

	newMoview = Movie{MovieName: name, ReleaseYear: year, MovieRating: rating}
	m.MovieList = append(m.MovieList, newMoview)
	err = m.Storage.Save(m.MovieList)
	if err != nil {
		return newMoview, err
	}

	return newMoview, err
}

// SortByRating - sort movie by rating
func (m *MovieRepo) SortByRating() {
	movies := m.MovieList

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].MovieRating > movies[j].MovieRating
	})

	m.MovieList = movies
}

// FindByYear - find movies of a year
func (m *MovieRepo) FindByYear(year int) []Movie {
	var movies []Movie

	for _, movie := range m.MovieList {
		if releaseYear, _ := strconv.Atoi(movie.ReleaseYear); releaseYear == year {
			movies = append(movies, movie)
		}
	}

	return movies
}

// FindByTitle - find movie with a keyword search
func (m *MovieRepo) FindByTitle(name string) []Movie {
	var movies []Movie
	re := regexp.MustCompile("(?i)" + name)

	for _, movie := range m.MovieList {
		if re.Match([]byte(movie.MovieName)) {
			movies = append(movies, movie)
		}
	}

	return movies
}
