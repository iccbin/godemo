package code

import (
	"html/template"
	"os"
	"strings"

)

func CreateHtml(rootPath string, modelInfo ModelInfo) error {

	path := rootPath + "/" + strings.ToLower(modelInfo.Name) + "s"
	os.Mkdir(rootPath+"/"+strings.ToLower(modelInfo.Name)+"s", 0777)

	file, err := os.OpenFile(path+"/"+"index.html", os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		return err
	}

	//template.FuncMap{"ToSlice" : ToSlice}
	errTmpl := tmpl.ExecuteTemplate(file, "index.tmpl", modelInfo)
	if err != nil {
		return errTmpl
	}

	return nil
}

//func CreateAll(rootPath string) error {
//	tmpl,err  := template.ParseGlob("./*")
//	if err != nil {
//		return err
//	}
//	files := make([]string, 0, 30)
//	filepath.Walk("./templates/", func(filename string, fi os.FileInfo, err error) {
//
//	})
//	errTmpl := tmpl.ExecuteTemplate(file)
//	return nil
//}

//func ToSlice(s string) string {
//	return strings.ToLower(s) + "s"
//}