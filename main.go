package main

import (
	"fmt"

	"unittest/hello"
)

func main() {

	//Sample1
	s1 := hello.GetHello1("DDD")
	fmt.Printf("%v", s1)

	//改行
	fmt.Println("")

	//Sample2
	h2 := hello.Hello{Say: "DDD"}
	s2 := h2.GetHello2()
	fmt.Printf("%v", s2)
}
