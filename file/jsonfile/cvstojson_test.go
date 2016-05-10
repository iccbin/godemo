package jsonfile

import (
	"testing"
	"fmt"
)

func TestCsvToJson(t *testing.T) {
		err := CsvToJson("/home/iven/Desktop/gotest.csv","/home/iven/Desktop/gotest.json")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
}

func BenchmarkCsvToJson(b *testing.B) {

	CsvToJson("/home/iven/Desktop/gotest.csv","/home/iven/Desktop/gotest.json")

}