package main

import "fmt"

func main() {
	var count int
	var totalMiles int
	var totalGallon int
	for count != 1 {
		fmt.Println("Enter Miles Driven")
		var miles int
		fmt.Scan(&miles)
		totalMiles += miles

		fmt.Println("Enter Gallon")
		var gallon int
		fmt.Scan(&gallon)
		totalGallon += gallon

		fmt.Println("Enter 1 to stop")
		fmt.Scan(&count)
	}
	var totalAmount float64 = float64(totalMiles / totalGallon)
	fmt.Println("your total amount is : ", totalAmount)
}
