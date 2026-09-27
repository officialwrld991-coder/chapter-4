package main

import "fmt"

func main() {
	fmt.Println("Enter a binary number")
	var binaryNumber int
	fmt.Scan(&binaryNumber)

	var sum int

	var count int = 1

	for binaryNumber > 0 {
		var lastDigit int = binaryNumber % 10
		binaryNumber = binaryNumber / 10

		sum += lastDigit * count
		count *= 2
	}

	fmt.Println(sum)
}
