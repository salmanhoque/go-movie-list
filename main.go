package main

import (
	"os"

	"github.com/salmanhoque/go-movie-list/cmd"
	"github.com/salmanhoque/go-movie-list/domain/movie"
	"github.com/salmanhoque/go-movie-list/storage"
)

const filename = "movies.json"

func main() {
	fileStorage := storage.JSONFileStorage{
		Filename: filename,
	}

	movieRepo := movie.Repo{Storage: fileStorage}

	cmd.Run(movieRepo, os.Args)
}
