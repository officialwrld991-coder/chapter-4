package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var number int
	fmt.Scan(&number)

	var sum int = 1

	for count := 1; count <= number; count++ {
		sum *= count
	}

	fmt.Println(sum)
}
