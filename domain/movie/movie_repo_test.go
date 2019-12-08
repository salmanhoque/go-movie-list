package movie

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (ms *MockStorage) Save(list interface{}) error {
	args := ms.Called()
	return args.Error(0)
}

func (ms *MockStorage) Read(list interface{}) error {
	args := ms.Called()
	return args.Error(0)
}

var (
	r   Repo
	err error
)

func TestAddToAddMovieToTheList(t *testing.T) {
	storage := new(MockStorage)
	r = Repo{Storage: storage}
	storage.On("Save").Return(err)

	movieName := "End Game"
	releaseYear := "2019"
	rating := 9.0
	expectedValue := Schema{movieName, releaseYear, rating}

	movie, err := r.Add(movieName, releaseYear, rating)

	assert.Equal(t, movie, expectedValue)
	assert.Nil(t, err)
}

func TestAddToAppendMoveAtTheEnd(t *testing.T) {
	storage := new(MockStorage)
	r = Repo{Storage: storage}
	storage.On("Save").Return(err)

	r.MovieList = append(r.MovieList, Schema{"Infinity War", "2018", 9.5})
	r.Add("End Game", "2019", 9.0)
	totalMovieCount := len(r.MovieList)

	assert.Equal(t, totalMovieCount, 2)
}

func TestAddToReturnErrorWhenFieldIsNotValid(t *testing.T) {
	storage := new(MockStorage)
	r := Repo{
		Storage: storage,
	}
	expecteErr := errors.New("Something went wrong")
	storage.On("Save").Return(expecteErr)

	_, err := r.Add("End Game", "2019", 9.0)
	assert.Equal(t, err, expecteErr)
}

func TestSortByRatingToReturnSortedList(t *testing.T) {
	hanSolo := Schema{"Han Solo", "2018", 6.5}
	endGame := Schema{"End Game", "2019", 9.2}
	spiderMan := Schema{"Spider Man", "2017", 8.2}
	movies := []Schema{hanSolo, endGame, spiderMan}
	s := new(MockStorage)

	expectedMoviesList := []Schema{endGame, spiderMan, hanSolo}
	r := Repo{
		MovieList: movies,
		Storage:   s,
	}

	r.SortByRating()

	assert.Equal(t, r.MovieList, expectedMoviesList)
}

func TestFindByYearToReturnMoviesOfGivenYear(t *testing.T) {
	s := new(MockStorage)
	hanSolo := Schema{"Han Solo", "2018", 6.5}
	infinityWar := Schema{"Infinity War", "2018", 9.2}
	spiderMan := Schema{"Spider Man", "2017", 8.2}
	movies := []Schema{hanSolo, infinityWar, spiderMan}

	expectedMoviesList := []Schema{hanSolo, infinityWar}
	r := Repo{
		MovieList: movies,
		Storage:   s,
	}

	returnedMoviesList := r.FindByYear(2018)

	assert.Equal(t, returnedMoviesList, expectedMoviesList)
}

func TestFindByNameToReturnMoviesOfGivenKeyword(t *testing.T) {
	s := new(MockStorage)
	hanSolo := Schema{"Han Solo", "2018", 6.5}
	infinityWar := Schema{"Infinity War", "2018", 9.2}
	spiderMan := Schema{"Spider Man", "2017", 8.2}

	r := Repo{
		MovieList: []Schema{hanSolo, infinityWar, spiderMan},
		Storage:   s,
	}
	expectedMovies := []Schema{hanSolo}

	assert.Equal(t, r.FindByTitle("Solo"), expectedMovies)
}

func TestFindByNameToReturnMoviesOfGivenKeywordInLowerCase(t *testing.T) {
	s := new(MockStorage)
	hanSolo := Schema{"Han Solo", "2018", 6.5}
	infinityWar := Schema{"Infinity War", "2018", 9.2}
	spiderMan := Schema{"Spider Man", "2017", 8.2}

	r := Repo{
		MovieList: []Schema{hanSolo, infinityWar, spiderMan},
		Storage:   s,
	}
	expectedMovies := []Schema{hanSolo}

	assert.Equal(t, r.FindByTitle("solo"), expectedMovies)
}
