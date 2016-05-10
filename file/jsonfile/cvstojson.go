package jsonfile

import (

	"bufio"
	"os"
	"io"
	"strings"
	"encoding/json"
	"sync"
)

type Student struct {
	Id string		`json:"id"`
	Name string		`json:"name"`
	Age string		`json:"age"`

}

type Tool struct {
	Mu sync.Mutex
	Wait sync.WaitGroup
	Buffer chan string
	Signal chan int
	Rows int
	ReadFile *os.File
	WriteFile *os.File
	End bool
}



func CsvToJson(readPath string,writePath string) error {
	tool := new(Tool)
	tool.Rows = 100
	tool.Buffer = make(chan string, tool.Rows)
	tool.Signal = make(chan int)
	tool.Wait.Add(1)
	var readErr,writeErr error
	tool.ReadFile,readErr = os.Open(readPath)
	tool.WriteFile,writeErr = os.OpenFile(writePath,os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)

	if readErr != nil {
		return readErr
	}
	if writeErr != nil {
		return writeErr
	}
	go ReadToFile(tool)
	//go Notify(tool)
	tool.Wait.Wait()
	return nil
}

func ReadToFile(tool *Tool) {
	reader := bufio.NewReader(tool.ReadFile)
	for i := 1; ; i++ {
		line,err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				tool.End = true
				close(tool.Buffer)
				//tool.Signal <- 1

				WriteToFile(tool)

				break
			}
		} else {

			if line != "" {
				//tool.Mu.Lock()
				tool.Buffer <- line
				//tool.Mu.Unlock()
				if i == tool.Rows {
					//tool.Signal <- 1
					WriteToFile(tool)
					i = 1
				}
			}
		}

	}
}



func WriteToFile(tool *Tool) {
	for i := 0; i < tool.Rows-1 ; i ++{
		line := <-tool.Buffer

		if line == "" {
			break;
		}
		lineSlice := strings.Split(line,",")
		if len(lineSlice) == 4 {
			st := &Student {
				lineSlice[0],
				lineSlice[1],
				lineSlice[2],
			}
			b,err := json.Marshal(st)
			tool.WriteFile.Write(b)
			tool.WriteFile.Write([]byte{'\n'})
			if err != nil {
				//return err

			}
		} else {
			break
		}
	}

	if tool.End {
		println(tool.End)
		tool.ReadFile.Close()
		tool.WriteFile.Close()

		tool.Wait.Done()
	}
}

func Notify(tool *Tool) {

	for {
		select {
		case val := <-tool.Signal:
			if val == 1 {
				tool.Mu.Lock()
				WriteToFile(tool)
				tool.Mu.Unlock()
			}

		}

	}
}
