package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var s []map[string]string
	var m map[string]string
	m = make(map[string]string)
	if s == nil {
		fmt.Println("s is nil")
	}
	if m == nil {
		fmt.Println("m is nil")
	}
	fmt.Println(s)
	fmt.Println(m)
	t1 := new(T)
	t2 :=  T{}
	var d float32 = 3.2
	fmt.Println(reflect.TypeOf(t1),",",reflect.TypeOf(t2))
	fmt.Println(reflect.ValueOf(t1),",",reflect.ValueOf(t2))
	fmt.Println(reflect.TypeOf(d), ",",reflect.ValueOf(d))
	v1 := reflect.ValueOf(t1)
	v2 := reflect.ValueOf(t1).Elem()
	if v2.CanSet() {
		fmt.Println("v2 can set")
	}
	fmt.Println(v1.Interface())
	fmt.Println(reflect.TypeOf(v2),reflect.ValueOf(v2))
	for i := 0; i < v2.Type().NumField(); i++ {
		fmt.Println(v2.Field(i).Type(),v2.Type().Field(i).Name,v2.Field(i).Interface())
		fmt.Println(v2.Type().Field(i).Type)
		m[v2.Type().Field(i).Name] = v2.Field(i).Type().String()
	}
	fmt.Println(m)
	for key,value := range m  {
		fmt.Println(key,value)
	}
	s=append(s[:5],s[5+1:]...)
}
