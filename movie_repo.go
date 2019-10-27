package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const fileName = "movies.json"

type movieRepo []movie

func (m movieRepo) listMovies() {
	fmt.Println()
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range m {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.MovieName, movie.ReleaseYear, movie.MovieRating)
	}

	fmt.Println()
}

func (m movieRepo) saveMovies() {
	moviesFromFile := m.readMoviesFromFile()
	m = append(m, moviesFromFile...)

	jsonData, _ := json.MarshalIndent(m, "", "  ")
	err := ioutil.WriteFile(fileName, jsonData, 0644)
	checkError("Can't write to file", err)

	println("Your movies saved successfully!")
}

func (m movieRepo) readMoviesFromFile() movieRepo {
	data, err := ioutil.ReadFile(fileName)
	checkError("Can't read the file", err)

	var moviesFromFile movieRepo
	err = json.Unmarshal(data, &moviesFromFile)
	checkError("Can't read all", err)

	return moviesFromFile
}

func (m movieRepo) listMoviesFromFile() {
	movies := m.readMoviesFromFile()

	movies.listMovies()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
