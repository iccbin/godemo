package main

import (
	"log"
	"os"
	"net/http"
	"godemo/web/controllers"
)

var logger *log.Logger



func main() {
	errLog := InitLog()
	if errLog != nil {
		logger.Fatalln(errLog.Error())

	}
	http.HandleFunc("/",controllers.Index)
	http.HandleFunc("/hello",controllers.Hello)
	http.HandleFunc("/login",controllers.Login)
	http.ListenAndServe(":9090", nil)

}



func InitLog() error  {
	file, err := os.OpenFile("./log.log",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return err;
	}
	logger = log.New(file,"",log.LstdFlags|log.Llongfile)
	return nil;
}
