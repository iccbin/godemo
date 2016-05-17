package createmodule

import "fmt"

func CreateStruct(controllerName string) string {
	var codeString string
	fmt.Sprintf("type %s struct {\n", controllerName)

	return codeString;
}