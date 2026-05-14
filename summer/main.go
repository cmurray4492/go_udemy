package main

import "fmt"

func main() {
	var sum int
	var num1 int

	fmt.Println("Please enter a number to sum: ")
	fmt.Scanln(&num1)

	for i := 1; i <= num1; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
