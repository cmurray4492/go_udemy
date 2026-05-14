package main

import "fmt"

func main() {
	var number int
	var ptr *int

	number = 20

	fmt.Println("Memory Address: ", &number)
	fmt.Println("Variable Value: ", number)

	ptr = &number

	fmt.Println("Memory Address: ", ptr)
	fmt.Println("Variable Value: ", *ptr)

	number = 10

	fmt.Println("Memory Address: ", ptr)
	fmt.Println("Variable Value: ", *ptr)

	*ptr = 5

	fmt.Println("Memory Address: ", ptr)
	fmt.Println("Variable Value: ", *ptr)

}
