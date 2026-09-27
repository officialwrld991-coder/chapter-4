package main

import "fmt"

func main() {
	fmt.Println("Enter firstNumber")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Println("Enter secondNumber")
	var secondNumber int
	fmt.Scan(&secondNumber)

	if firstNumber == secondNumber {
		fmt.Println(0)
	} else if firstNumber > secondNumber {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}

}
