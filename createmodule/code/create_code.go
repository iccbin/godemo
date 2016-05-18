package code

import (
	"io/ioutil"
	//"os"
	"strings"
	"fmt"
	"text/template"
	"os"
)


func CreateFile(modelInfo ModelInfo) error {


	tmpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		return err
	}

	files,err := ListFile("./templates/", ".tmpl")
	if err != nil {
		return err
	}
	//template.FuncMap{"ToSlice" : ToSlice}


	for _,fileName := range files {
		fmt.Println(fileName)
		file, _ := os.OpenFile("./product/"+ fileName, os.O_CREATE | os.O_RDWR, 0666)
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
	//PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			//files = append(files, dirPth+PthSep+fi.Name())
			files = append(files, fi.Name())
			fmt.Println(fi.Name())
		}
	}
	return files, nil
}
