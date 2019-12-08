package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// JSONFileStorage - a data type that implements persistance interface.
type JSONFileStorage struct {
	Filename string
}

// Save - saves struct as JSON
func (s JSONFileStorage) Save(list interface{}) error {
	var err error
	var jsonData []byte
	jsonData, err = json.MarshalIndent(list, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Unable to encode to JSON")
	}

	f, err := os.OpenFile(s.Filename, os.O_CREATE|os.O_WRONLY, 0644)
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

// Read - reads a JSON file
func (s JSONFileStorage) Read(list interface{}) error {
	var data []byte
	var err error

	if _, err = os.Stat(s.Filename); os.IsNotExist(err) {
		if _, err = os.Create(s.Filename); err != nil {
			return errors.Wrap(err, "Could not able to read from the file!")
		}

		return err
	}

	if data, err = ioutil.ReadFile(s.Filename); err != nil {
		return errors.Wrap(err, "Could not able to read from the file!")
	}

	if err = json.Unmarshal(data, &list); err != nil {
		return errors.Wrap(err, "Could not able to parse the file")
	}

	return err
}
