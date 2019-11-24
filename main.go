package main

import (
	domain "github.com/salmanhoque/go-movie-list/domain/movie"
	storage "github.com/salmanhoque/go-movie-list/storage"
	"os"
)

const fileName = "movies.json"

func main() {
	var fileStorage storage.JSONFileStorage

	mr := domain.MovieRepo{Storage: fileStorage}

	run(mr, os.Args)
}
