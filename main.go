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

type movies []movie

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
	var movies movies
	fmt.Println(help)

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
			fmt.Println(help)
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
		if err != nil {
			log.Fatal("Can't write to file", err)
		}
	}

	println("Your movies list saved successfully!")
}

func (m *movies) readMoviesFromFile() {
	b, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal("Can't read the file", err)
	}

	r := csv.NewReader(strings.NewReader(string(b)))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Can't read all", err)
	}

	for _, r := range records {
		rating, _ := strconv.ParseFloat(r[2], 64)

		om := movie{
			movieName:   r[0],
			releaseYear: r[1],
			movieRating: rating,
		}

		*m = append(*m, om)
	}

	fmt.Println("\n Got your movies!")
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
