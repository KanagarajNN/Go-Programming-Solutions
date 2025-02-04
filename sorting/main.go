// main.go - Entry point to run Bubble Sort in Go
//
// Description:
// - This file is the entry point for running the Bubble Sort algorithm.
// - It initializes an array, sorts it using BubbleSort, and prints the output.
//
// Usage:
// - Run the program using:
//     go run main.go
//
// Metadata:
// - Author: Kanagaraj N N
// - Date: February 1, 2025
// - Version: 1.0
// - License: MIT

package main

import (
	"fmt"

	"github.com/kanagarajnn/go-programming-solutions/sorting/bubble_sort"
)

func main() {
	// Sample array to be sorted
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted array:", arr)

	// Call BubbleSort function
	bubble_sort.BubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
