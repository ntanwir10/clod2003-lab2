package main

// Palindrome function to check if a string is a palindrome
func Palindrome(str string) bool {
   lastIdx := len(str) - 1
   // using for loop
   for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
      if str[i] != str[lastIdx-i] {
         return false
      }
   }
   return true
}
