package main

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON File Storage", func() {

	Describe("save", func() {
		const testFileName = "test_fixture/save_test_movies.json"

		It("saves json data to a file", func() {
			var storage jsonFileStorage
			movies := []movie{
				{"End Game", "2018", 9.2},
				{"Infinity War", "2019", 9.0},
			}

			Expect(storage.save(movies, testFileName)).Should(Succeed())

			_, err := os.Stat(testFileName)
			Expect(os.IsNotExist(err)).Should(BeFalse())

			os.Remove(testFileName)
		})
	})

	Describe("read", func() {
		const testFileName = "test_fixture/test_movie_data.json"

		It("reads json data from a file", func() {
			var storage jsonFileStorage
			var movies []movie

			Expect(storage.read(&movies, testFileName)).Should(Succeed())

			Expect(movies[0]).Should(Equal(movie{"End Game", "2018", 9.2}))
		})
	})

})
