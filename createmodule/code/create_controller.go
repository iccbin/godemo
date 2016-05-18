package code

import (

	"fmt"
	"os"
	"html/template"
	"strings"
)


func CreateController(path string, modelInfo ModelInfo) error {

	f,err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("no such file or directory")
	}
	if !f.IsDir() {
		return fmt.Errorf("path is not directory")
	}

	filePath := fmt.Sprintf("%s/%s_controller.go",path,strings.ToLower(modelInfo.Name))

	file,err := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	tmpl,err := template.ParseFiles("controllers.tmpl")
	if err != nil {
		return err
	}

	errL := tmpl.ExecuteTemplate(file, "controllers.tmpl", modelInfo)
	if err != nil {
		return errL
	}
	return nil
}

