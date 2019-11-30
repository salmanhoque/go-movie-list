package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/salmanhoque/go-movie-list/domain/movie"
)

// Run - stars the movie app in command line
func Run(mr movie.Repo, args []string) {
	err := mr.Storage.Read(&mr.MovieList)
	if err != nil {
		printError(err)
	}

	if len(args) > 1 {
		switch args[1] {
		case "list":
			listMovies(mr.MovieList)
		case "add":
			addMovie(mr, args)
		case "list-by-rating":
			mr.SortByRating()
			listMovies(mr.MovieList)
		case "find-by-year":
			findByYear(mr, args)
		case "find-by-title":
			findByTitle(mr, args)
		default:
			help()
		}
	} else {
		help()
	}
}

func help() {
	fmt.Println()
	fmt.Printf("%-30s%-30s\n", "Command", "Usage")
	fmt.Printf("%-30s%-30s\n", "list", "Show all movies")
	fmt.Printf("%-30s%-30s\n", "add",
		"add new movie. add --name [movie name] --year [release year] --rating [your rating]")
	fmt.Printf("%-30s%-30s\n", "list-by-rating",
		"List all movies sorted by rating")
	fmt.Printf("%-30s%-30s\n", "find-by-year",
		"List all movies of a year, find-by-year --year 2019")
	fmt.Printf("%-30s%-30s\n", "find-by-title",
		"Find a movie by a keyword, find-by-title --keyword game")
	fmt.Printf("%-30s%-30s\n", "help", "show all commands")
	fmt.Println()
}

func askQuestion(question string) string {
	r := bufio.NewScanner(os.Stdin)
	fmt.Printf(question + " ")
	r.Scan()

	return r.Text()
}

func addMovie(mr movie.Repo, args []string) {
	addCommand := flag.NewFlagSet("add", flag.ContinueOnError)
	name := addCommand.String("name", "", "Movie name")
	year := addCommand.String("year", "", "Release year")
	rating := addCommand.Float64("rating", 0.0, "Your rating")

	addCommand.Parse(args[2:])

	if *name == "" || *year == "" || *rating == 0.0 {
		println("\nadd new movie. add --name [movie name] --year [release year] --rating [your rating]\n")
		addCommand.Usage()
		return
	}

	m, err := mr.Add(*name, *year, *rating)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("\nAdded %s(%s) with a rating %.2f\n\n", m.MovieName, m.ReleaseYear, m.MovieRating)
}

func findByYear(mr movie.Repo, args []string) {
	findByYearCommand := flag.NewFlagSet("find-by-year", flag.ContinueOnError)
	year := findByYearCommand.Int("year", 0, "Enter a year to filter movies")

	findByYearCommand.Parse(args[2:])

	movies := mr.FindByYear(*year)
	listMovies(movies)
}

func findByTitle(mr movie.Repo, args []string) {
	findByYearCommand := flag.NewFlagSet("find-by-title", flag.ContinueOnError)
	keyword := findByYearCommand.String("keyword", "", "Enter a keyword to filter movies")

	findByYearCommand.Parse(args[2:])

	movies := mr.FindByTitle(*keyword)
	listMovies(movies)
}

func listMovies(m []movie.Schema) {
	fmt.Println()
	fmt.Printf("|%-30s|%-30s|%-30s|\n", "Movie Name", "Rlease Year", "Rating")

	for _, movie := range m {
		fmt.Printf("|%-30s|%-30s|%-30.2f|\n",
			movie.MovieName, movie.ReleaseYear, movie.MovieRating)
	}

	fmt.Println()
}

func printError(err error) {
	fmt.Printf("\n\n=== Error: %+v\n\n", err)
}
