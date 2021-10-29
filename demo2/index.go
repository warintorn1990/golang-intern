package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("a")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}
