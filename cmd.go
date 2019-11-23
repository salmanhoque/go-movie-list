package main

// https://blog.neillyons.io/mocking-command-line-flags-and-stdin-in-go/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func run(mr movieRepo, args []string) {
	err := mr.storage.read(&mr.movieList, fileName)
	if err != nil {
		printError(err)
	}

	if len(args) > 1 {
		switch args[1] {
		case "list":
			listMovies(mr)
		case "add":
			addMovie(mr, args)
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
	fmt.Printf("%-30s%-30s\n", "list-by-rating", "List all movies sorted by rating")
	fmt.Printf("%-30s%-30s\n", "find-by-year", "List all movies of a year")
	fmt.Printf("%-30s%-30s\n", "find-by-name", "Find a movie by name")
	fmt.Printf("%-30s%-30s\n", "help", "show all commands")
	fmt.Println()
}

func askQuestion(question string) string {
	r := bufio.NewScanner(os.Stdin)
	fmt.Printf(question + " ")
	r.Scan()

	return r.Text()
}

func addMovie(mr movieRepo, args []string) {
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

	m, err := mr.add(*name, *year, *rating)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("\nAdded %s(%s) with a rating %.2f\n\n", m.MovieName, m.ReleaseYear, m.MovieRating)
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