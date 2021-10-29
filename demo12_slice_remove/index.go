package main

import "fmt"

func main() {
	// var number = make([]int, 3, 5)
	var number = []int{1, 2, 3}
	// var number []int
	showSlice(number)

	// //remove
	// number = number[1:len(number)]
	// showSlice(number)

	// //traling remove
	// number = number[0 : len(number)-1]
	// showSlice(number)

	//remove Index
	removeIndex(number, 1)
	showSlice(number)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
