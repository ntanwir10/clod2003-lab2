package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group B's Week 4 Project!")

	num1 := 56
	num2 := 35
	num3 := 20
	max := findMax(num1, num2, num3)
	fmt.Printf("The maximum of %d, %d, and %d is %d\n", num1, num2, num3, max)
}
