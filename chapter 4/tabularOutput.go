package main

import "fmt"

func main() {
	for count := 0; count <= 5; count++ {
		fmt.Printf("%d %d %d %d \n", count, count*count, count*count*count, count*count*count*count)
	}
}
