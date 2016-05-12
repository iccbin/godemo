package jsonfilev2

import (
	"testing"
	"fmt"
	"os"
	"sync"
)

func TestCsvToJson(t *testing.T) {

}

func BenchmarkCsvToJson(b *testing.B) {

	var size int64
	var readPath string = "/home/iven/Desktop/test.txt"
	var writePath string = "/home/iven/Desktop/test1.txt"

	file,err := os.Open(readPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	size = fileInfo.Size()
	file.Close()

	fmt.Println("size:", size)
	littleSize := size/6
	var begin int64 = 0

	mu := sync.Mutex{}
	wait := new(sync.WaitGroup)
	wait.Add(6)
	for i := 0; i < 6; i++ {
		mu.Lock()
		go WriteFile(begin, begin + littleSize, readPath, writePath, wait)
		begin += littleSize
		mu.Unlock()

	}
	wait.Wait()

}