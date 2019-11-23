package main

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie App")
}

type MockStorage struct {
	mock.Mock
}

func (ms *MockStorage) save(list interface{}, fileName string) error {
	args := ms.Called()
	return args.Error(0)
}

func (ms *MockStorage) read(list interface{}, fileName string) error {
	args := ms.Called()
	return args.Error(0)
}

var _ = Describe("Movie Repo", func() {

	Describe("add", func() {
		Context("when movie fields are valid", func() {
			var (
				m   movieRepo
				err error
			)

			BeforeEach(func() {
				storage := new(MockStorage)
				m = movieRepo{storage: storage}
				storage.On("save").Return(err)
			})

			It("adds movie to the list", func() {
				movieName := "End Game"
				releaseYear := "2019"
				rating := 9.0

				expectedValue := movie{movieName, releaseYear, rating}

				Expect(m.add(movieName, releaseYear, rating)).To(Equal(expectedValue))
			})

			It("append movie to the list", func() {
				m.movieList = append(m.movieList, movie{"Infinity War", "2018", 9.5})

				m.add("End Game", "2019", 9.0)

				Expect(len(m.movieList)).To(Equal(2))
			})
		})

		Context("when movie fields are not valid", func() {
			It("returns an error message", func() {
				storage := new(MockStorage)
				m := movieRepo{
					storage: storage,
				}
				expecteErr := errors.New("Something went wrong")
				storage.On("save").Return(expecteErr)

				_, actual := m.add("End Game", "2019", 9.0)
				Expect(actual).Should(MatchError(expecteErr))
			})
		})
	})

	Describe("sortByRating", func() {
		It("returns movies sorted by rating", func() {
			hasSolo := movie{"Han Solo", "2018", 6.5}
			endGame := movie{"End Game", "2019", 9.2}
			spiderMan := movie{"Spider Man", "2017", 8.2}

			movies := []movie{hasSolo, endGame, spiderMan}
			var storage jsonFileStorage
			sortedMovies := []movie{endGame, spiderMan, hasSolo}

			mr := movieRepo{
				movieList: movies,
				storage:   storage,
			}

			mr.sortByRating()

			Expect(mr.movieList).Should(Equal(sortedMovies))
		})
	})
})
