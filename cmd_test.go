package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	domain "github.com/salmanhoque/go-movie-list/domain/movie"
	"github.com/stretchr/testify/mock"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie command-line app")
}

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

var _ = Describe("Main Interface", func() {
	var (
		mr  domain.MovieRepo
		err error
	)

	BeforeEach(func() {
		storage := new(MockStorage)
		movies := []domain.Movie{
			{"End Game", "2018", 9.2},
			{"Infinity War", "2019", 9.0},
		}

		mr = domain.MovieRepo{
			MovieList: movies,
			Storage:   storage,
		}

		storage.On("Read").Return(err)
		storage.On("Save").Return(err)
	})

	Describe("list", func() {
		It("lists all movies", func() {
			args := []string{"filename", "list"}

			output := captureStdout(func() {
				run(mr, args)
			})

			Expect(output).Should(MatchRegexp("End Game"))
			Expect(output).Should(MatchRegexp("Infinity War"))
		})
	})

	Describe("add", func() {
		It("add movies", func() {
			args := []string{"filename", "add", "--name", "Joker", "--year", "2019", "--rating", "9.2"}

			output := captureStdout(func() {
				run(mr, args)
			})

			Expect(strings.TrimSpace(output)).Should(MatchRegexp("Added Joker"))
		})
	})

	Describe("list-by-rating", func() {
		It("sorts movies by rating", func() {
			args := []string{"filename", "list-by-rating"}

			output := captureStdout(func() {
				run(mr, args)
			})

			Expect(strings.TrimSpace(output)).Should(MatchRegexp("End Game"))
		})
	})

	Describe("find-by-year", func() {
		It("finds movies of a year", func() {
			args := []string{"filename", "find-by-year", "--year", "2019"}

			output := captureStdout(func() {
				run(mr, args)
			})

			Expect(strings.TrimSpace(output)).Should(MatchRegexp("Infinity War"))
		})
	})

	Describe("find-by-title", func() {
		It("finds movies by a keyword", func() {
			args := []string{"filename", "find-by-title", "--keyword", "game"}

			output := captureStdout(func() {
				run(mr, args)
			})

			Expect(strings.TrimSpace(output)).Should(MatchRegexp("End Game"))
		})
	})
})
