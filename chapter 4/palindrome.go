package main

import "fmt"

func main() {
	fmt.Println("Enter digit")
	var digit string
	fmt.Scan(&digit)

	var reversedDigit string
	for count := len(digit) - 1; count >= 0; count-- {
		reversedDigit = reversedDigit + string(digit[count])
	}

	if digit == reversedDigit {
		fmt.Println("it is a palindrome")
	} else {
		fmt.Println("it is not a palindrome")
	}
}
