package main

import (
	"fmt"
	"os"

	"godemo/createmodule/example/config"
	"godemo/createmodule/example/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	router.Any("/", users.Index)
	router.GET("/users-new", users.New)
	router.POST("/users", users.Create)
	router.DELETE("/users/:id", users.Delete)
	router.GET("/users-edit/:id", users.Edit)
	router.PUT("/users/:id", users.Update)
	router.GET("/users/:id", users.Show)
	router.Run(":3001")
}
