package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// Persistence is used to save and read from file or database
type Persistence interface {
	save(list interface{}, fileName string) error
	read(list interface{}, fileName string) error
}

const fileName = "movies.json"

var fileStorage jsonFileStorage
var movieR movieRepo = movieRepo{storage: fileStorage}

func main() {
	err := movieR.storage.read(&movieR.movieList, fileName)
	if err != nil {
		printError(err)
	}

	help()

	for {
		action := prompt()

		switch action {
		case "add":
			addMovie()
		case "list":
			listMovies(movieR)
		case "help":
			help()
		case "exit":
			fmt.Println("\nBye Bye :)")
			os.Exit(0)
		default:
			prompt()
		}
	}
}

func prompt() string {
	return askQuestion("What you would like to do:")
}

func help() {
	var help = `
Welcome to "Go Movie Go", You one stop movie directory.

Here are some of the command you can do using this great movie app!

add:              Add new movie name with year and rating
list:             List all movies
list-by-rating:   List all movies sorted by rating
find-by-year:     List all movies of a year
find-by-name:     Find a movie by name
help:             See all commands 
exit:             Exit out from this app.
`

	fmt.Println(help)
}

func askQuestion(question string) string {
	r := bufio.NewScanner(os.Stdin)
	fmt.Printf(question + " ")
	r.Scan()

	return r.Text()
}

func addMovie() {
	name := askQuestion("Enter movie name:")
	year := askQuestion("Enter release year:")

	ratingStr := askQuestion("Rating:")
	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		printError(errors.Wrap(err, "Rating should be a number"))
		return
	}

	m, err := movieR.add(name, year, rating)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("\nAdded %s(%s) with a rating %.2f\n\n",
		m.MovieName, m.ReleaseYear, m.MovieRating)
}

func listMovies(mr movieRepo) {
	fmt.Println()
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range mr.movieList {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.MovieName, movie.ReleaseYear, movie.MovieRating)
	}

	fmt.Println()
}

func printError(err error) {
	fmt.Printf("\n\n=== Error: %+v\n\n", err)
}
