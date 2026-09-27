package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var number int
	fmt.Scan(&number)

	for count := 1; count <= number; count++ {
		for counter := 1; counter <= count; counter++ {
			fmt.Print("x")
		}
		fmt.Println()
	}
}
