package domain

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie Domain Logic")
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

var _ = Describe("Movie Repo", func() {

	Describe("add", func() {
		Context("when movie fields are valid", func() {
			var (
				m   MovieRepo
				err error
			)

			BeforeEach(func() {
				storage := new(MockStorage)
				m = MovieRepo{Storage: storage}
				storage.On("Save").Return(err)
			})

			It("adds movie to the list", func() {
				movieName := "End Game"
				releaseYear := "2019"
				rating := 9.0

				expectedValue := Movie{movieName, releaseYear, rating}

				Expect(m.Add(movieName, releaseYear, rating)).To(Equal(expectedValue))
			})

			It("append movie to the list", func() {
				m.MovieList = append(m.MovieList, Movie{"Infinity War", "2018", 9.5})

				m.Add("End Game", "2019", 9.0)

				Expect(len(m.MovieList)).To(Equal(2))
			})
		})

		Context("when movie fields are not valid", func() {
			It("returns an error message", func() {
				storage := new(MockStorage)
				m := MovieRepo{
					Storage: storage,
				}
				expecteErr := errors.New("Something went wrong")
				storage.On("Save").Return(expecteErr)

				_, actual := m.Add("End Game", "2019", 9.0)
				Expect(actual).Should(MatchError(expecteErr))
			})
		})
	})

	Describe("sortByRating", func() {
		It("returns movies sorted by rating", func() {
			hanSolo := Movie{"Han Solo", "2018", 6.5}
			endGame := Movie{"End Game", "2019", 9.2}
			spiderMan := Movie{"Spider Man", "2017", 8.2}

			movies := []Movie{hanSolo, endGame, spiderMan}
			s := new(MockStorage)
			sortedMovies := []Movie{endGame, spiderMan, hanSolo}

			mr := MovieRepo{
				MovieList: movies,
				Storage:   s,
			}

			mr.SortByRating()

			Expect(mr.MovieList).Should(Equal(sortedMovies))
		})
	})

	Describe("findByYear", func() {
		It("returns movies sorted by rating", func() {
			s := new(MockStorage)
			hanSolo := Movie{"Han Solo", "2018", 6.5}
			infinityWar := Movie{"Infinity War", "2018", 9.2}
			spiderMan := Movie{"Spider Man", "2017", 8.2}

			movies := []Movie{hanSolo, infinityWar, spiderMan}
			moviesInYear2018 := []Movie{hanSolo, infinityWar}

			mr := MovieRepo{
				MovieList: movies,
				Storage:   s,
			}

			actual := mr.FindByYear(2018)

			Expect(actual).Should(Equal(moviesInYear2018))
		})
	})

	Describe("findByName", func() {
		s := new(MockStorage)
		hanSolo := Movie{"Han Solo", "2018", 6.5}
		infinityWar := Movie{"Infinity War", "2018", 9.2}
		spiderMan := Movie{"Spider Man", "2017", 8.2}

		movies := []Movie{hanSolo, infinityWar, spiderMan}

		mr := MovieRepo{
			MovieList: movies,
			Storage:   s,
		}

		Context("when search keyword is in titlecase", func() {
			It("returns matched movies", func() {
				actual := mr.FindByTitle("Solo")
				expected := []Movie{hanSolo}

				Expect(actual).Should(Equal(expected))
			})
		})

		Context("when search keyword is in lowercase", func() {
			It("returns matched movies", func() {
				actual := mr.FindByTitle("solo")
				expected := []Movie{hanSolo}

				Expect(actual).Should(Equal(expected))
			})
		})
	})
})
