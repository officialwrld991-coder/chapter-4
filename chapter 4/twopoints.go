package main

import "fmt"

func main() {
	fmt.Println("Enter x1")
	var x1 string
	fmt.Scan(&x1)

	fmt.Println("Enter y1")
	var y1 string
	fmt.Scan(&y1)

	fmt.Println("Enter x2")
	var x2 string
	fmt.Scan(&x2)

	fmt.Println("Enter y2")
	var y2 string
	fmt.Scan(&y2)

	if x1 == x2 {
		fmt.Println("x axis")
	}
	if y1 == y2 {
		fmt.Println("y axis")
	}
}
