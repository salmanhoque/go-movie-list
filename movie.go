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
	addReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter movie name: ")
	text, _ := addReader.ReadString('\n')
	m.movieName = strings.TrimSpace(text)

	fmt.Println("Enter release year: ")
	text, _ = addReader.ReadString('\n')
	m.releaseYear = strings.TrimSpace(text)

	fmt.Println("Rating: ")
	text, _ = addReader.ReadString('\n')
	m.movieRating, _ = strconv.ParseFloat(strings.TrimSpace(text), 64)

	// Print added movie
	t := fmt.Sprintf("Added \"%s\"(%s) with a rating %.2f\n",
		m.movieName, m.releaseYear, m.movieRating)
	fmt.Println(t)

	return m
}
