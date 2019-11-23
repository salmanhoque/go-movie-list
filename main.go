package main

import "os"

// Persistence is used to save and read from file or database
type Persistence interface {
	save(list interface{}, fileName string) error
	read(list interface{}, fileName string) error
}

const fileName = "movies.json"

func main() {
	var fileStorage jsonFileStorage
	var mr movieRepo = movieRepo{storage: fileStorage}

	run(mr, os.Args)
}
