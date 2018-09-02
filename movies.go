package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type movies []movie

func (m movies) listMovies() {
	fmt.Print("\n")
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range m {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.movieName, movie.releaseYear, movie.movieRating)
	}

	fmt.Print("\n")
}

func (m movies) saveMovies() {
	addReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a filename: ")
	filename, _ := addReader.ReadString('\n')

	file, err := os.Create(strings.TrimSpace(filename) + ".csv")

	if err != nil {
		log.Fatal("Can't create file", err)
	}

	write := csv.NewWriter(file)
	defer write.Flush()

	for _, movie := range m {
		ratingToStr := strconv.FormatFloat(movie.movieRating, 'f', 2, 64)
		data := []string{movie.movieName, movie.releaseYear, ratingToStr}
		err := write.Write(data)
		checkError("Can't write to file", err)
	}

	println("Your movies saved successfully!")
}

func (m *movies) readMoviesFromFile() {
	addReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your filename: ")
	filename, _ := addReader.ReadString('\n')

	b, err := ioutil.ReadFile(strings.TrimSpace(filename) + ".csv")
	checkError("Can't read the file", err)

	r := csv.NewReader(strings.NewReader(string(b)))

	records, err := r.ReadAll()
	checkError("Can't read all", err)

	for _, r := range records {
		rating, _ := strconv.ParseFloat(r[2], 64)

		om := movie{
			movieName:   r[0],
			releaseYear: r[1],
			movieRating: rating,
		}

		*m = append(*m, om)
	}

	m.listMovies()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
