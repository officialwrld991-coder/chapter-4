package main

import "fmt"

func main() {
	var basePay float32 = 200.0

	var count int

	var salesAmount float32

	var totalAmount float32

	for count != 1 {
		fmt.Println("Enter sale")
		fmt.Scan(&salesAmount)
		totalAmount += salesAmount

		fmt.Println("Enter 1 to stop")
		fmt.Scan(&count)
	}

	var total = totalAmount * 0.09
	fmt.Println("Sales commission is:", total+basePay)
}
