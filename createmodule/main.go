package main

import (
	"io/ioutil"
	"strings"
	"text/template"
	"os"
	"fmt"
	"reflect"


	"godemo/createmodule/example/models"

	"github.com/chuckpreslar/inflect"
)

type ModelInfo struct {
	Name string
	FieldMap map[string]string
}

func CreateFile(modelInfo ModelInfo) error {

	funcs := template.FuncMap{
		"title" : strings.Title,
		"pluralize" : inflect.Pluralize,
		"lower" : strings.ToLower,
	}

	tmpl, err := template.New("").Funcs(funcs).ParseGlob("./templates/*")
	if err != nil {
		return err
	}


	files,err := ListFile("./templates/", ".tmpl")
	if err != nil {
		return err
	}

	for _,fileName := range files {

		var newFileName string
		if strings.Contains(fileName, "controller") {
			newFileName = strings.TrimSuffix(fileName, "tmpl") + "go"
		} else {
			newFileName = strings.TrimSuffix(fileName, "tmpl") + "html"
		}

		file, err := os.OpenFile("./product/"+ newFileName, os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
			return err
		}


		errTmpl := tmpl.ExecuteTemplate(file, fileName, modelInfo)
		if err != nil {
			return errTmpl
		}
	}

	return nil
}


func ListFile(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {

			files = append(files, fi.Name())
		}
	}
	return files, nil
}



func createCode()  {

	modelInfo,err := ReflectModel(&models.User{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	split := strings.Split( reflect.TypeOf(models.User{}).String(), ".")
	modelInfo.Name = strings.ToLower(split[1])

	createErr := CreateFile(modelInfo)
	if createErr != nil {
		fmt.Println(createErr.Error())
	}
}


func ReflectModel(i interface{}) (ctrl ModelInfo, err error) {

	info := ModelInfo{}

	val := reflect.ValueOf(i).Elem()

	if val.Kind() != reflect.Struct {
		return info, fmt.Errorf("interface kind is not struct")
	}

	info.FieldMap = make(map[string]string)
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		info.FieldMap[typeField.Name] = typeField.Type.String()
	}

	return info, nil
}

func DealString(arg string) string {
	return strings.ToLower(arg)
}

func main()  {
	createCode()
}