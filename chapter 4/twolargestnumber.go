package main

import "fmt"

func main() {
	var count int
	var largest int = 1
	var secondlargest int = 0
	for count = 1; count <= 10; count++ {
		fmt.Println("Enter number")
		var number int
		fmt.Scan(&number)

		if number > largest {
			secondlargest = largest
			largest = number
		}

		if number > secondlargest && number < largest {
			secondlargest = number
		}

	}
	fmt.Println("Largest is: ", largest)
	fmt.Println("secondlargest is: ", secondlargest)
}
