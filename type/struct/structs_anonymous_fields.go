package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
	aa func()
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
	innerS //anonymous field

}

type handler func(s string, t string)

func Tt(aa func(s string,t string))  {
	aa("1", "2")
}

func Ss(handler handler){
	handler("3","4")
}

func main() {
	inner := new(innerS)

	inner.aa = func() {
		fmt.Println("aa func")
	}
	inner.aa()

	Tt(func(s string, t string) {
		fmt.Println(s,t)
	})

	outer := &outerS{}
	fmt.Println(outer)
}
