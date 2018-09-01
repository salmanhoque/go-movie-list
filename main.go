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

type movie struct {
	movieName   string
	releaseYear string
	movieRating float64
}

var help = `
Welcome to "Go Movie Go", You one stop movie directory.

Here are some of the command you can do using this great movie app!

add	
 - Add new movie name with year and rating
list
 - List all movies
list-by-rating
 - List all movies sorted by rating
save
 - Save movie to a file
read
 - Load your movies from a file
help
 - See all commands 
exit
 - Exit out from this app.
`

func main() {
	var movies []movie
	fmt.Println(help)

	for {
		action := prompt()

		if action == "add" {
			m := addMovie()
			movies = append(movies, m)
		}

		if action == "list" {
			listMovies(movies)
		}

		if action == "save" {
			saveMovies(movies)
		}

		if action == "read" {
			movies = readMoviesFromFile()
		}

		if action == "exit" {
			fmt.Println("\nBye Bye :)")
			break
		}

		if action == "help" {
			fmt.Println(help)
		}
	}
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What you would like to do: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func listMovies(movies []movie) {
	fmt.Print("\n")
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range movies {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.movieName, movie.releaseYear, movie.movieRating)
	}

	fmt.Print("\n")
}

func saveMovies(movies []movie) {
	addReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a filename: ")
	filename, _ := addReader.ReadString('\n')

	file, err := os.Create(strings.TrimSpace(filename) + ".csv")

	if err != nil {
		log.Fatal("Can't create file", err)
	}

	write := csv.NewWriter(file)
	defer write.Flush()

	for _, movie := range movies {
		ratingToStr := strconv.FormatFloat(movie.movieRating, 'f', 2, 64)
		data := []string{movie.movieName, movie.releaseYear, ratingToStr}
		err := write.Write(data)
		if err != nil {
			log.Fatal("Can't write to file", err)
		}
	}

	println("Your movies list saved successfully!")
}

func readMoviesFromFile() []movie {
	var movies []movie

	b, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal("Can't read the file", err)
	}

	r := csv.NewReader(strings.NewReader(string(b)))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Can't read all", err)
	}

	for _, m := range records {
		rating, _ := strconv.ParseFloat(m[2], 64)

		m := movie{
			movieName:   m[0],
			releaseYear: m[1],
			movieRating: rating,
		}

		movies = append(movies, m)
	}

	fmt.Println("\n Got your movies!")
	return movies
}

func addMovie() movie {
	var m movie
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
