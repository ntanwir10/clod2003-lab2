package main

import (
	"fmt"
)

func main() {
	var fahrenheit float32
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius := fahrenheitToCelsius(fahrenheit)
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)
}
