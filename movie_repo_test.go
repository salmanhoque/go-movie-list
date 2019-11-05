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

func (ms *MockStorage) save(m []movie, fileName string) error {
	args := ms.Called()
	return args.Error(0)
}

func (ms *MockStorage) read(m *[]movie, fileName string) error {
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

})
