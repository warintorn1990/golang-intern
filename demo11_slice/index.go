package main

import "fmt"

func main() {
	// var number = make([]int, 3, 5)
	var number []int
	number = append(number, 1)
	showSlice(number)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
