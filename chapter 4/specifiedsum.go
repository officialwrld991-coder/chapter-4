package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scan(&input)

	var firstInput int = input

	var count int

	var sum int

	for count == 0 {
		fmt.Println("Enter a number")
		var digit int
		fmt.Scan(&digit)

		sum += digit

		if sum > firstInput || sum == firstInput {
			break
		}
	}
}
