// Language: Go
// Solution 1: Bubble Sort Algorithm 
// Time: Worst - O(N^2), Avg - O(N^2), Best - O(N)
// Space: O(1)

package main

import "fmt"

func bubbleSort(array []int) []int {
	isSorted := false
	counter := 0

	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-counter-1; i++ {
			if array[i] > array[i+1] {
				swapElements(i, i+1, array)
				isSorted = false
			}
		}
		counter++
	}

	return array
}

func swapElements(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	fmt.Println(array)

	// Sort array using Bubble Sort
	sortedArray := bubbleSort(array)

	fmt.Println("Sorted array:")
	fmt.Println(sortedArray)
}
