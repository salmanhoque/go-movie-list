package main

import (
	"bytes"
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

var _ = Describe("Main Interface", func() {
	var (
		mr  movieRepo
		err error
	)

	BeforeEach(func() {
		storage := new(MockStorage)
		movies := []movie{
			{"End Game", "2018", 9.2},
			{"Infinity War", "2019", 9.0},
		}

		mr = movieRepo{
			movieList: movies,
			storage:   storage,
		}

		storage.On("read").Return(err)
		storage.On("save").Return(err)
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
})