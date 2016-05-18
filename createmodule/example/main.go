package main

import (
	"fmt"
	"os"

	"godemo/createmodule/example/config"
	"godemo/createmodule/example/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/users/*")

	mysqlConfig := config.Mysql{DataBase: "study", UserName: "root", Password: "123456", Parameters: "charset=utf8&parseTime=True&loc=Local"}
	DB, err := gorm.Open("mysql", mysqlConfig.String())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	application := controllers.ApplicationController{DB: DB}
	users := controllers.UsersController{application}

	router.GET("/", users.Index)

	router.Run(":3001")
}
