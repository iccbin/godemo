package code

import (
	"io/ioutil"
	"strings"
	"text/template"
	"os"
)

func CreateFile(modelInfo ModelInfo) error {


	tmpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		return err
	}
	tmpl,err = tmpl.Funcs(template.FuncMap{"DealString":DealString})
	files,err := ListFile("./templates/", ".tmpl")
	if err != nil {
		return err
	}

	for _,fileName := range files {

		var newFileName string
		if strings.Contains(fileName, "controller") {
			newFileName = strings.TrimSuffix(fileName, "tmpl") + "go"
		} else {
			modelInfo.Name = strings.ToLower(modelInfo.Name) + "s"
			newFileName = strings.TrimSuffix(fileName, "tmpl") + "html"
		}

		file, err := os.OpenFile("./product/"+ newFileName, os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
			return err
		}

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

	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件

			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func DealString(arg string) string {
	return strings.ToLower(arg)
}