package main

import _ package1

func Foo() (string, error){
}

func main() {
	Foo()
	s, _ := Foo()

	ch := make(chan int)
	<-ch

	var x interface{} = 7
	i := x.(int) // type assertion

	var Map = make(map[string]string)
	v := Map["key"]

	for k := range Map {
	}

	for _, v := range Map {
	}
}