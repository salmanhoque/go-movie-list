package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var movies movies
	help()

	for {
		action := prompt()

		switch action {
		case "add":
			var m movie
			movies = append(movies, m.addMovie())
		case "list":
			movies.listMovies()
		case "save":
			movies.saveMovies()
		case "read":
			movies.readMoviesFromFile()
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What you would like to do: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func help() {
	var help = `
Welcome to "Go Movie Go", You one stop movie directory.

Here are some of the command you can do using this great movie app!

add:              Add new movie name with year and rating
list:             List all movies
list-by-rating:   List all movies sorted by rating
save:             Save movie to a file
read:             Load your movies from a file
help:             See all commands 
exit:             Exit out from this app.
`

	fmt.Println(help)
}
