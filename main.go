package main

import "fmt"

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
}
