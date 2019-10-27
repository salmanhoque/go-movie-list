package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const fileName = "movies.json"

type movieRepo []movie

func (m *movieRepo) add(name string, year string, rating float64) movie {
	newMoview := movie{MovieName: name, ReleaseYear: year, MovieRating: rating}
	*m = append(*m, newMoview)
	m.save()

	return newMoview
}

func (m *movieRepo) save() {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	checkError("Unable to create/file the file", err)

	jsonData, err := json.MarshalIndent(m, "", "  ")
	checkError("Unable to encode to JSON", err)
	_, err = f.Write(jsonData)
	checkError("Unable to write to the file", err)

	defer f.Close()
}

func (m *movieRepo) all() movieRepo {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return *m
	}

	data, err := ioutil.ReadFile(fileName)
	checkError("Can't read the file", err)

	err = json.Unmarshal(data, &m)
	checkError("Can't read all", err)

	return *m
}
