package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// a data type that implements the interface
type jsonFileStorage struct{}

func (s jsonFileStorage) save(m []movie, fileName string) error {
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

func (s jsonFileStorage) read(m *[]movie, fileName string) error {
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
