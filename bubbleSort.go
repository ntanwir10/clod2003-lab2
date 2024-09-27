// Created by Neelotpal Chaulia

package main

import "fmt"

// BubbleSort function to sort an array
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func ExampleBubblesort() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array:", arr)

    BubbleSort(arr)

    fmt.Println("Sorted array:", arr)
}
