package createmodule

import (
	"reflect"
	"fmt"
)



func ReflectModel(i interface{}) error {
	val := reflect.ValueOf(i).Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("fsdfsdfafdsfdsa")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		fmt.Printf("name:%s  type:%s value:%v\n", typeField.Name, typeField.Type.String(), val.Field(i).Interface())
	}

	return nil
}

