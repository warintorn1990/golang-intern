package main

import "fmt"

func main() {
	msg := "Some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)

	changeMessage(&msg)
	fmt.Println(msg)

}

func changeMessage(aPointer *string) {
	*aPointer = "NEW MESSAGE"
}
