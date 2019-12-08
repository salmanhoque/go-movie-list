package movie

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
				r   Repo
				err error
			)

			BeforeEach(func() {
				storage := new(MockStorage)
				r = Repo{Storage: storage}
				storage.On("Save").Return(err)
			})

			It("adds movie to the list", func() {
				movieName := "End Game"
				releaseYear := "2019"
				rating := 9.0

				expectedValue := Schema{movieName, releaseYear, rating}

				Expect(r.Add(movieName, releaseYear, rating)).To(Equal(expectedValue))
			})

			It("append movie to the list", func() {
				r.MovieList = append(r.MovieList, Schema{"Infinity War", "2018", 9.5})

				r.Add("End Game", "2019", 9.0)

				Expect(len(r.MovieList)).To(Equal(2))
			})
		})

		Context("when movie fields are not valid", func() {
			It("returns an error message", func() {
				storage := new(MockStorage)
				r := Repo{
					Storage: storage,
				}
				expecteErr := errors.New("Something went wrong")
				storage.On("Save").Return(expecteErr)

				_, actual := r.Add("End Game", "2019", 9.0)
				Expect(actual).Should(MatchError(expecteErr))
			})
		})
	})

	Describe("sortByRating", func() {
		It("returns movies sorted by rating", func() {
			hanSolo := Schema{"Han Solo", "2018", 6.5}
			endGame := Schema{"End Game", "2019", 9.2}
			spiderMan := Schema{"Spider Man", "2017", 8.2}

			movies := []Schema{hanSolo, endGame, spiderMan}
			s := new(MockStorage)
			sortedMovies := []Schema{endGame, spiderMan, hanSolo}

			r := Repo{
				MovieList: movies,
				Storage:   s,
			}

			r.SortByRating()

			Expect(r.MovieList).Should(Equal(sortedMovies))
		})
	})

	Describe("findByYear", func() {
		It("returns movies sorted by rating", func() {
			s := new(MockStorage)
			hanSolo := Schema{"Han Solo", "2018", 6.5}
			infinityWar := Schema{"Infinity War", "2018", 9.2}
			spiderMan := Schema{"Spider Man", "2017", 8.2}

			movies := []Schema{hanSolo, infinityWar, spiderMan}
			moviesInYear2018 := []Schema{hanSolo, infinityWar}

			r := Repo{
				MovieList: movies,
				Storage:   s,
			}

			actual := r.FindByYear(2018)

			Expect(actual).Should(Equal(moviesInYear2018))
		})
	})

	Describe("findByName", func() {
		s := new(MockStorage)
		hanSolo := Schema{"Han Solo", "2018", 6.5}
		infinityWar := Schema{"Infinity War", "2018", 9.2}
		spiderMan := Schema{"Spider Man", "2017", 8.2}

		movies := []Schema{hanSolo, infinityWar, spiderMan}

		r := Repo{
			MovieList: movies,
			Storage:   s,
		}

		Context("when search keyword is in titlecase", func() {
			It("returns matched movies", func() {
				actual := r.FindByTitle("Solo")
				expected := []Schema{hanSolo}

				Expect(actual).Should(Equal(expected))
			})
		})

		Context("when search keyword is in lowercase", func() {
			It("returns matched movies", func() {
				actual := r.FindByTitle("solo")
				expected := []Schema{hanSolo}

				Expect(actual).Should(Equal(expected))
			})
		})
	})
})
