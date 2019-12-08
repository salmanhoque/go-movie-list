package storage

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storage system")
}

var _ = Describe("JSON File Storage", func() {
	type movie struct {
		Title  string
		Year   string
		Rating float64
	}

	var storage JSONFileStorage

	Describe("save", func() {
		It("saves json data to a file", func() {
			const testFileName = "test_fixture/save_movies_data.json"

			storage = JSONFileStorage{
				Filename: testFileName,
			}

			movies := []movie{
				{"End Game", "2018", 9.2},
				{"Infinity War", "2019", 9.0},
			}

			Expect(storage.Save(movies)).Should(Succeed())

			_, err := os.Stat(testFileName)
			Expect(os.IsNotExist(err)).Should(BeFalse())

			os.Remove(testFileName)
		})
	})

	Describe("read", func() {
		It("reads json data from a file", func() {
			const testFileName = "test_fixture/read_movies_data.json"

			storage = JSONFileStorage{
				Filename: testFileName,
			}

			var movies []movie

			Expect(storage.Read(&movies)).Should(Succeed())

			Expect(movies[0]).Should(Equal(movie{"End Game", "2018", 9.2}))
		})
	})

	It("creates file if not present", func() {
		storage = JSONFileStorage{
			Filename: "test_fixture/missing-file.json",
		}

		var movies []movie

		Expect(storage.Read(&movies)).Should(Succeed())

		// clean up after test
		os.Remove(storage.Filename)
	})
})
