![Go CI](https://github.com/kanagarajnn/go-programming-solutions/actions/workflows/ci-go.yml/badge.svg)

# Bubble Sort Implementation in Go

## Description
This project implements the **Bubble Sort** algorithm in Go. Bubble Sort is a simple comparison-based sorting algorithm that repeatedly iterates through an array, swapping adjacent elements if they are in the wrong order. The algorithm continues iterating until no more swaps are needed, meaning the array is sorted.

## Features
- Implements Bubble Sort algorithm in Go
- Handles various edge cases like empty arrays, already sorted arrays, and negative numbers
- Unit tested using Go's `testing` package
- Structured as a **Go module** for easy dependency management and testing

## Time Complexity
- **Worst Case:** \(O(N^2)\) - When the array is reverse sorted
- **Average Case:** \(O(N^2)\) - For a randomly shuffled array
- **Best Case:** \(O(N)\) - When the array is already sorted

## Space Complexity
- \(O(1)\) (In-place sorting algorithm)

## Directory Structure
```
go-programming-solutions/
  ├── sorting/
  │   ├── bubble_sort/
  │   │   ├── bubble_sort.go
  │   │   ├── bubble_sort_test.go
  │   ├── main.go
  ├── go.mod
  ├── README.md
```

## Installation & Execution
Ensure you have **Go 1.20+** installed.

### **1. Clone the repository**
```bash
git clone https://github.com/your-username/go-programming-solutions.git
cd go-programming-solutions
```

### **2. Initialize the Go module**
```bash
go mod tidy
```

### **3. Run the program**
```bash
go run sorting/main.go
```

## Running Unit Tests
This project includes **Go tests** to validate the correctness of the Bubble Sort algorithm.

### **Run tests using Go**
```bash
go test ./sorting/bubble_sort -v
```

## Example Usage
```go
arr := []int{64, 34, 25, 12, 22, 11, 90}
fmt.Println("Unsorted array:", arr)

bubble_sort.BubbleSort(arr)

fmt.Println("Sorted array:", arr)
```
### **Expected Output:**
```
Unsorted array: [64 34 25 12 22 11 90]
Sorted array: [11 12 22 25 34 64 90]
```

## Continuous Integration (CI) with GitHub Actions
This project is configured with GitHub Actions for automated testing.

## CI/CD Workflow Details:
- Runs tests automatically on every push and pull request to the main branch.
- Uses Go 1.20 on an Ubuntu runner.
- Ensures code correctness before merging.

## How to Check CI/CD Status:
- Push your changes:
```sh
$ git push origin main
```
- Navigate to your GitHub repository → Actions tab.
- Check the workflow **Go CI** to view test results.

## License
This project is licensed under the **MIT License**.

## Metadata
- Author: Kanagaraj N N
- Email: n.n.kanagaraj.upm@gmail.com
- GitHub: https://github.com/kanagarajnn

