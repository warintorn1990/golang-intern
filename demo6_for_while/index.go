package main

import "fmt"

func main() {
	fnFor()
	forWhile()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("For index %d\n", index)
	}
}

func forWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Printf("For While index %d\n", index)
	}
}

func forWhileUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		index++
		fmt.Printf("For While index %d\n", index)
	}
}
