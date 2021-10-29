package main

import "fmt"

func main() {
	sameValue := 10
	if sameValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if sameValue == 10 || sameValue < 20 {
		fmt.Println("sameValue == 10 ||  sameValue < 20")
	} else {
		fmt.Println("NOT sameValue == 10 ||  sameValue < 20")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
	fnSwitchcase()
}

func doSomething() string {
	return "ok"
}

func fnSwitchcase() {
	index := 2
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("default")
		break
	}
}
