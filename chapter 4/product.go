package main

import "fmt"

var product int = 5
var x int = 5

func main() {
	product *= product
	x++
	fmt.Println(product)
	fmt.Println(x)
}
