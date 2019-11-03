package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const fileName = "movies.json"

type movieRepo []movie

func (m *movieRepo) add(name string, year string, rating float64) (movie, error) {
	var err error
	var newMoview movie

	newMoview = movie{MovieName: name, ReleaseYear: year, MovieRating: rating}
	*m = append(*m, newMoview)
	err = m.save()
	if err != nil {
		return newMoview, err
	}

	return newMoview, err
}

func (m *movieRepo) save() error {
	var err error
	var jsonData []byte
	jsonData, err = json.MarshalIndent(m, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Unable to encode to JSON")
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return errors.Wrap(err, "Unable to create a file")
	}

	_, err = f.Write(jsonData)
	if err != nil {
		return errors.Wrap(err, "Unable to write to the file")
	}

	return err
}

func (m *movieRepo) all() error {
	var data []byte
	var err error

	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		return errors.Wrap(err, "Can't able to find the file!")
	}

	if data, err = ioutil.ReadFile(fileName); err != nil {
		return errors.Wrap(err, "Can't able to read the file!")
	}

	if err = json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, "Can't able to parse the file")
	}

	return err
}
