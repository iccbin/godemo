package main

import (
	"fmt"
	"os"

	"godemo/createmodule/example/config"

	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

type User struct  {
	gorm.Model
	Name string `gorm:"size:10"`
	Num int

}

type Language struct {
	ID int
	Name string `gorm:"index:idx_name_code`
	Code string `gorm:"index:idx_name_code`
}
func main() {
	mysqlConfig := config.Mysql{DataBase: "gorm", UserName: "root", Password: "123456", Parameters: "charset=utf8&parseTime=True&loc=Local"}
	db, err := gorm.Open("mysql", mysqlConfig.String())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&User{},&Language{})
	db.Close()
}


