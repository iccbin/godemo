package createmodule

import (
	"testing"

	"godemo/createmodule/example/models"
	"fmt"
)

func TestReflectModel(t *testing.T) {
	err := ReflectModel(&models.User{})
	if err != nil {
		fmt.Println(err.Error())
	}
}