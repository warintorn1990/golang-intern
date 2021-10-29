package main

import "fmt"

func main() {
	fn1()
	fn2("Good moreing")
	fn3("Good moreing", 1)
	const a, b = 1, 2
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))

	var x, y int = swap(a, b)
	fmt.Printf("%d,%d => %d,%d\n", a, b, x, y)

	x, y = swap2(10, 20)
	fmt.Printf("%d,%d => %d,%d\n", 10, 20, x, y)
}

func fn1() {
	fmt.Println("Arm")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Println(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x, y int) {
	y = a
	x = b
	return
}
