package main

import (
	domain "github.com/salmanhoque/go-movie-list/domain/movie"
	storage "github.com/salmanhoque/go-movie-list/storage"
	"os"
)

const filename = "movies.json"

func main() {
	fileStorage := storage.JSONFileStorage{
		Filename: filename,
	}

	mr := domain.MovieRepo{Storage: fileStorage}

	run(mr, os.Args)
}
