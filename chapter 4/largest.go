package main

import "fmt"

func main() {
	var count int
	var largest int = 1
	for count = 1; count <= 10; count++ {
		fmt.Println("Enter number")
		var number int
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}
	}
	fmt.Println("Largest is: ", largest)
}
