package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Group B's Week 4 Project!")
  
  // Function - 1: Palindrome function by Srinidhi Sivakumar
	fmt.Println("Function 1: Palindrome function submitted by Srinidhi Sivakumar ")
   
   // declare and initialize the string
   str := "madam"
   fmt.Println("Check the given word is palindrome or not,\n Given Word =",str)
 
   // calling the function
   if Palindrome(str) {
      fmt.Printf("'%s' is a palindrome\n", str)
   } else {
      fmt.Printf("'%s' is not a palindrome\n", str)
   }  
  
  // Function - 2: Maximum of 3 Numbers by Anjani
	num1 := 56
	num2 := 35
	num3 := 20
	max := findMax(num1, num2, num3)
	fmt.Printf("The maximum of %d, %d, and %d is %d\n", num1, num2, num3, max)  

   // Function - 3: Fahrenheit to Celsius by Nauman Tanwir
	var fahrenheit float32
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius := fahrenheitToCelsius(fahrenheit)
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)
}
