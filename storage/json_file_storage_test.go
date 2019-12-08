package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type movie struct {
	Title  string
	Year   string
	Rating float64
}

var storage JSONFileStorage

func TestSaveToSaveAJSONFile(t *testing.T) {
	const testFileName = "test_fixture/save_movies_data.json"

	storage = JSONFileStorage{
		Filename: testFileName,
	}
	movies := []movie{
		{"End Game", "2018", 9.2},
		{"Infinity War", "2019", 9.0},
	}

	err := storage.Save(movies)
	assert.Nil(t, err)
	assert.FileExists(t, testFileName)

	os.Remove(testFileName)
}

func TestReadToReadFromAJSONFile(t *testing.T) {
	const testFileName = "test_fixture/read_movies_data.json"
	var movies []movie

	storage = JSONFileStorage{
		Filename: testFileName,
	}

	err := storage.Read(&movies)
	assert.Nil(t, err)
	assert.Equal(t, movies[0], movie{"End Game", "2018", 9.2})
}

func TestReadToCreateAFileIfNotExists(t *testing.T) {
	var movies []movie

	storage = JSONFileStorage{
		Filename: "test_fixture/missing-file.json",
	}

	err := storage.Read(&movies)
	assert.Nil(t, err)

	os.Remove(storage.Filename)
}
