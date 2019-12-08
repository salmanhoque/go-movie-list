package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/salmanhoque/go-movie-list/domain/movie"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

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

type CmdTestSuite struct {
	suite.Suite
	repo movie.Repo
	err  error
}

func (suite *CmdTestSuite) SetupTest() {
	storage := new(MockStorage)
	movies := []movie.Schema{
		{MovieName: "End Game", ReleaseYear: "2018", MovieRating: 9.2},
		{MovieName: "Infinity War", ReleaseYear: "2019", MovieRating: 9.0},
	}

	suite.repo = movie.Repo{
		MovieList: movies,
		Storage:   storage,
	}

	storage.On("Read").Return(suite.err)
	storage.On("Save").Return(suite.err)
}

func (suite *CmdTestSuite) TestListToListAllMovies() {
	args := []string{"filename", "list"}

	output := captureStdout(func() {
		Run(suite.repo, args)
	})

	suite.Contains(output, "End Game")
	suite.Contains(output, "Infinity War")
}

func (suite *CmdTestSuite) TestAddToAddMoviesToTheList() {
	args := []string{"filename", "add", "--name", "Joker", "--year", "2019", "--rating", "9.2"}

	output := captureStdout(func() {
		Run(suite.repo, args)
	})

	suite.Contains(output, "Added Joker")
}

func (suite *CmdTestSuite) TestListByRatingReturnsFilteredMoviesByRating() {
	args := []string{"filename", "list-by-rating"}

	output := captureStdout(func() {
		Run(suite.repo, args)
	})

	suite.Contains(output, "End Game")
}

func (suite *CmdTestSuite) TestfindByYearReturnsFilteredMoviesByYear() {
	args := []string{"filename", "find-by-year", "--year", "2019"}

	output := captureStdout(func() {
		Run(suite.repo, args)
	})

	suite.Contains(output, "Infinity War")
}

func (suite *CmdTestSuite) TestFindByTitleReturnsFilteredMoviesByKeywor() {
	args := []string{"filename", "find-by-title", "--keyword", "game"}

	output := captureStdout(func() {
		Run(suite.repo, args)
	})

	suite.Contains(output, "End Game")
}

func TestCmdTestSuite(t *testing.T) {
	suite.Run(t, new(CmdTestSuite))
}
