package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type movie struct {
	movieName   string
	releaseYear string
	movieRating float64
}

func (m movie) addMovie() movie {
	m.movieName = movieName()
	m.releaseYear = releaseYear()
	m.movieRating = movieRating()

	// Print added movie
	t := fmt.Sprintf("Added \"%s\"(%s) with a rating %.2f\n",
		m.movieName, m.releaseYear, m.movieRating)
	fmt.Println(t)

	return m
}

func movieName() string {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter movie name: ")
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

func releaseYear() string {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter release year: ")
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

func movieRating() float64 {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Rating: ")
	text, _ := r.ReadString('\n')
	rating, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)

	return rating
}
