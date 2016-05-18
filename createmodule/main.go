package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"reflect"

	"godemo/createmodule/example/models"
	"godemo/createmodule/code"
)


func main()  {

	//modelInfo,err := code.ReflectModel(&models.Temp{}, "godemo/createmodule/example/models")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//modelInfo.ModelName = reflect.TypeOf(models.Temp{}).String()
	//split := strings.Split(modelInfo.ModelName, ".")
	//modelInfo.Name = split[1]
	//modelInfo.VarName = strings.ToLower(modelInfo.Name) + "s"
	//
	//ctrlErr := code.CreateController( "example/controllers", modelInfo)
	//if ctrlErr != nil {
	//	fmt.Println(ctrlErr.Error())
	//}
	//
	//htmlErr := code.CreateHtml("example/templates", modelInfo)
	//if htmlErr != nil {
	//	fmt.Println(htmlErr.Error())
	//}
	createCode()

}

func createCode()  {
	modelInfo,err := code.ReflectModel(&models.Temp{}, "godemo/createmodule/example/models")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	modelInfo.ModelName = reflect.TypeOf(models.Temp{}).String()
	split := strings.Split(modelInfo.ModelName, ".")
	modelInfo.Name = split[1]
	modelInfo.VarName = strings.ToLower(modelInfo.Name) + "s"

	code.CreateFile(modelInfo)
}
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
			fmt.Println(fi.Name())
		}
	}
	return files, nil
}