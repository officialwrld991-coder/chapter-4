package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)

	fmt.Println("Enter your annual earning")
	var annualEarning float32
	fmt.Scan(&annualEarning)

	var totalTax float64

	if annualEarning <= 30000 {
		totalTax = 30000 / 0.15
		fmt.Printf("Tax rate for %s is %d", name, totalTax)
	} else {
		totalTax = 30000 / 0.20
		fmt.Printf("Tax rate for %s is %d", name, totalTax)
	}
}
