package jsonfilev2


import (
	"os"
	"fmt"
	"sync"
)





func WriteFile(begin int64, end int64, readPath string, writePath string, wait *sync.WaitGroup) error {
	fmt.Println("begin:", begin, "end:", end)


	readFile,err := os.Open(readPath)
	if err != nil {
		return err
	}
	defer readFile.Close()
	readFile.Seek(begin, 0)

	writeFile,err := os.OpenFile(writePath, os.O_RDWR, 0777)
	if err != nil {
		wait.Done()
		return err
	}
	defer writeFile.Close()
	writeFile.Seek(begin, 0)

	var size int64 = 1028
	if (end - begin) < size {
		size = end - begin
	}


	for i := begin; int64(i) < end; i += int64(size) {
		buffer := make([]byte, size)
		n, err := readFile.Read(buffer)
		if err != nil {
			wait.Done()
			return err
		}

		if int64(n) == size {
			writeFile.Write(buffer)
		} else {
			writeFile.Write(buffer[0:n])
		}



	}
	wait.Done()
	return nil
}