package main

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

func main() {
	var p1 product
	p1.name = "Android"
	p1.price = 100
	p1.stock = 20
	show(p1)
}

func show(p product) {
	fmt.Println(p)
}
