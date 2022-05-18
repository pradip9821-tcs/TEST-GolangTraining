package main

import "fmt"

func main() {
	
	// Normal for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()

	// Nested for loop
	for i := 0; i < 10; i++ {
		for j := i; j < i+1; j++ {
			fmt.Print(i, " ", j, "\t")
		}
	}
	fmt.Println()
}
