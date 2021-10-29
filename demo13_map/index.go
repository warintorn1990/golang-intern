package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers)

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	fmt.Println(colors)

	var course = make(map[string]map[string]int)

	// course["Android"] = map[string]int{"price": 200}

	course["Android"] = make(map[string]int)
	course["Android"]["price"] = 200
	course["Android"]["code"] = 0001

	fmt.Println(course)
}
