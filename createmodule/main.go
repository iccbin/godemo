package main

import (
	"fmt"

	"godemo/createmodule/example/models"
	"godemo/createmodule/code"

	"strings"
	"reflect"
)


func main()  {

	modelInfo,err := code.ReflectModel(&models.Temp{}, "godemo/createmodule/example/models")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	modelInfo.ModelName = reflect.TypeOf(models.Temp{}).String()
	split := strings.Split(modelInfo.ModelName, ".")
	modelInfo.Name = split[1]
	modelInfo.VarName = strings.ToLower(modelInfo.Name) + "s"

	ctrlErr := code.CreateController( "example/controllers", modelInfo)
	if ctrlErr != nil {
		fmt.Println(ctrlErr.Error())
	}

	htmlErr := code.CreateHtml("example/templates", modelInfo)
	if htmlErr != nil {
		fmt.Println(htmlErr.Error())
	}
}
