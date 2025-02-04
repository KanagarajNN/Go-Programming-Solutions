// bubble_sort.go - Implementation of Bubble Sort in Go
//
// Description:
// - This file implements the Bubble Sort algorithm in Go.
// - It sorts an array of integers in ascending order using the Bubble Sort technique.
//
// Time Complexity:
// - Worst Case: O(N^2) (When the array is reverse sorted)
// - Average Case: O(N^2) (For a randomly shuffled array)
// - Best Case: O(N) (When the array is already sorted)
//
// Space Complexity:
// - O(1) (In-place sorting algorithm)
//
// Usage:
// - Compile and run using:
//     go run main.go
//
// Metadata:
// - Author: Kanagaraj N N
// - Date: February 4, 2025
// - Version: 1.0
// - License: MIT

package bubble_sort

import (
	"fmt"
)

// BubbleSort sorts an array using the Bubble Sort algorithm.
func BubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	n := len(arr)
	swapped := false

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// PrintArray prints the array in a readable format.
func PrintArray(arr []int) {
	fmt.Println(arr)
}

// // main function for running the Bubble Sort
// func main() {
// 	arr := []int{64, 34, 25, 12, 22, 11, 90}
// 	fmt.Print("Unsorted array: ")
// 	PrintArray(arr)

// 	BubbleSort(arr)

// 	fmt.Print("Sorted array: ")
// 	PrintArray(arr)
// }
