package main

import "fmt"

func main() {
	cource := []string{"Android", "IOS", "React"}

	for index, item := range cource {
		fmt.Printf("%d . %s\n", index, item)
	}

	for _, item := range cource {
		fmt.Printf("%s\n", item)
	}
}
