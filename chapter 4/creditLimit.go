package main

import "fmt"

func main() {
	var accountNumber int
	var initialBalance int
	var totalItems int
	var totalCredits int
	var allowedCredits int

	fmt.Println("Enter Account Number")
	fmt.Scan(&accountNumber)

	fmt.Println("Enter Initial Balance")
	fmt.Scan(&initialBalance)

	fmt.Println("Enter Total Items")
	fmt.Scan(&totalItems)

	fmt.Println("Enter Total Credits")
	fmt.Scan(&totalCredits)

	fmt.Println("Enter Allowed Credits")
	fmt.Scan(&allowedCredits)

	var totalAmount int = initialBalance + totalItems - totalCredits

	if totalAmount > allowedCredits {
		fmt.Println("Amount Exceeded Limit")
	} else {
		fmt.Println("Amount Hasn't Exceeded Limit")
	}
}
