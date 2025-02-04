// bubble_sort_test.go - Unit Tests for Bubble Sort in Go
//
// Description:
// - This file contains unit tests for the Bubble Sort function using Go's `testing` package.
//
// Usage:
// - Run tests using:
//     go test ./sorting/bubble_sort -v
//
// Metadata:
// - Author: Kanagaraj N N
// - Date: February 4, 2025
// - Version: 1.0
// - License: MIT

package bubble_sort

import (
	"reflect"
	"testing"
)

// TestAlreadySorted tests an already sorted array.
func TestAlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}

// TestReverseSorted tests a reverse sorted array.
func TestReverseSorted(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}

// TestRandomOrder tests an array with random order.
func TestRandomOrder(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}

// TestSingleElement tests an array with a single element.
func TestSingleElement(t *testing.T) {
	arr := []int{42}
	expected := []int{42}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}

// TestEmptyArray tests an empty array.
func TestEmptyArray(t *testing.T) {
	arr := []int{}
	expected := []int{}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}

// TestNegativeNumbers tests an array with negative numbers.
func TestNegativeNumbers(t *testing.T) {
	arr := []int{10, -1, 3, 8, 5, 2, -3, 4}
	expected := []int{-3, -1, 2, 3, 4, 5, 8, 10}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort(%v) = %v, want %v", arr, arr, expected)
	}
}
