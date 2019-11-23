package main

import "sort"

type movieRepo struct {
	movieList []movie
	storage   Persistence
}

func (m *movieRepo) add(name string, year string, rating float64) (movie, error) {
	var err error
	var newMoview movie

	newMoview = movie{MovieName: name, ReleaseYear: year, MovieRating: rating}
	m.movieList = append(m.movieList, newMoview)
	err = m.storage.save(m.movieList, fileName)
	if err != nil {
		return newMoview, err
	}

	return newMoview, err
}

func (m *movieRepo) sortByRating() {
	movies := m.movieList

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].MovieRating > movies[j].MovieRating
	})

	m.movieList = movies
}
