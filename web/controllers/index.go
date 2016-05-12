package controllers

import (
	"fmt"
	"net/http"

	"html/template"
)

func Index(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println("url_long",r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key:",k,"======value:",v)
	}
	fmt.Fprint(w,"index")
}


func Hello(w http.ResponseWriter,r *http.Request){

	fmt.Fprint(w,"hello")
}


func Login(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		t,_ := template.ParseFiles("../views/login.html")
		t.Execute(w,nil)
	} else {
		fmt.Fprint(w,"login")
	}
}