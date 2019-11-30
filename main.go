package main

import (
	"os"

	cmd "github.com/salmanhoque/go-movie-list/cmd"
	domain "github.com/salmanhoque/go-movie-list/domain/movie"
	storage "github.com/salmanhoque/go-movie-list/storage"
)

const filename = "movies.json"

func main() {
	fileStorage := storage.JSONFileStorage{
		Filename: filename,
	}

	mr := domain.MovieRepo{Storage: fileStorage}

	cmd.Run(mr, os.Args)
}
