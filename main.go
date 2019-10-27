package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var movieR movieRepo

func main() {
	help()

	for {
		action := prompt()

		switch action {
		case "add":
			addMovie()
		case "list":
			listMovies()
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
	rating, _ := strconv.ParseFloat(ratingStr, 64)

	m := movieR.add(name, year, rating)

	fmt.Printf("\nAdded %s(%s) with a rating %.2f\n\n",
		m.MovieName, m.ReleaseYear, m.MovieRating)
}

func listMovies() {
	fmt.Println()
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range movieR.all() {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.MovieName, movie.ReleaseYear, movie.MovieRating)
	}

	fmt.Println()
}
