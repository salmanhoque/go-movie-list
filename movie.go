package main

import (
	"fmt"
	"strconv"
)

type movie struct {
	MovieName   string
	ReleaseYear string
	MovieRating float64
}

func (m movie) addMovie() movie {
	m.MovieName = askQuestion("Enter movie name:")
	m.ReleaseYear = askQuestion("Enter release year:")

	rating := askQuestion("Rating:")
	m.MovieRating, _ = strconv.ParseFloat(rating, 64)

	fmt.Printf("\nAdded %s(%s) with a rating %.2f\n\n",
		m.MovieName, m.ReleaseYear, m.MovieRating)

	return m
}
