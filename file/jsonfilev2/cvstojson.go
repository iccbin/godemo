package jsonfilev2

import (
	"os"
)




func CheckError(e error){
	if e != nil {
		panic(e)
	}
}

func CvsToJson(readPath string,writePath string){
	
}

func ReadFile(readPath string){
	f, err := os.Open(readPath)
	CheckError(err)
	f.
}

func WriteFile()  {
	
}