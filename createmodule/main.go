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
	ModelName string
	PluralizeTitle string
	PluralizeName string
	UnderscoreName string
}

func CreateFile(modelInfo ModelInfo) error {

	tmpl, err := template.New("").Delims("[[","]]").ParseGlob("./templates/*")
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





func ReflectModel(name string, i interface{}) (ctrl ModelInfo, err error) {

	info := ModelInfo{Name:name}
	info.ModelName = strings.Title(name)
	info.PluralizeTitle = strings.Title(inflect.Pluralize(name))
	info.PluralizeName = strings.ToLower(inflect.Pluralize(name))
	info.UnderscoreName = strings.ToLower(name)

	val := reflect.ValueOf(i).Elem()

	if val.Kind() != reflect.Struct {
		return info, fmt.Errorf("interface kind is not struct")
	}

	info.FieldMap = make(map[string]string)
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		//info.FieldMap[typeField.Name] = typeField.Type.String()
		info.FieldMap[typeField.Name] = strings.ToLower(typeField.Name)
	}

	return info, nil
}

func createCode(name string, i interface{})  {

	modelInfo,err := ReflectModel(name, i)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	modelInfo.Name = name

	createErr := CreateFile(modelInfo)
	if createErr != nil {
		fmt.Println(createErr.Error())
	}
}

func main()  {
	createCode("user", &models.User{})
}