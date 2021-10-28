package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Start")

	//Explicit declarlation
	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true
	const tmp4 int = 0

	//Implicit
	tmp5 := 0
	tmp6 := "OOP"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)
	fmt.Println(tmp5)
	fmt.Println(tmp6)
	run()
	run()
	run()
}

func run() {
	count++
	fmt.Println(count)
}
