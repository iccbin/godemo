package code

import (
	"fmt"
	"reflect"
)

type ModelInfo struct {
	ModelPath string
	Name string
	VarName string
	ModelName string
	FieldMap map[string]string
}

func ReflectModel(i interface{}, modelPath string) (ctrl ModelInfo, err error) {

	info := ModelInfo{ModelPath:modelPath}

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


