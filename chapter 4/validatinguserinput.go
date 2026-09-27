package main

import "fmt"

func main() {
	var count int = 0
	for count != 1 {
		fmt.Printf("Enter input")
		fmt.Scan(&count)
		if count == 2 {
			break
		}
	}
}
