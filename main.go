package main

import "fmt"

func main() {
	var rows int = 5 // Number of rows for the pyramid
	// fmt.Print("Enter number of rows: ")
	// fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {
		// Print leading spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		// New line after each row
		fmt.Println()
	}
}
