package main

import (
	"fmt"
)

func foo(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += 1
	}
	return m
}

func sumArr(data []int) int {

	total := 0

	for i := 0; i < len(data); i++ {
		total += data[i]
	}

	//Implement your solution here
	// fmt.Println(data, len(data))

	return total
}

func SequentialSearch(data []int, value int) bool {
	//Implement your solution here
	for i := 0; i < len(data); i++ {
		if data[i] == value {
			return true
		}
	}
	return false //Return false if value not found
}

func MaxSubArraySum(data []int) int {

	//Implement your solution here
	max := -99999999999999999
	// low := 0
	// high := len(data) - 1
	for i := 0; i < len(data)-1; i++ {
		sum := 0
		for j := i; j < len(data); j++ {
			sum += data[j]
			if max < sum {
				max = sum
				// low = i
				// high = j
			}
		}

	}

	return max
}

func RotateArray(data []int, k int) int {
	//Implement your solution here
	for j := 0; j < k; j++ {
		temp := data[0]
		for i := 0; i < len(data)-1; i++ {
			data[i] = data[i+1]

			if i+1 == len(data)-1 {
				data[i+1] = temp
			}

		}
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("data[i]", data[i])
	}

	return 0

}

func WaveArray(arr []int) {
	//Implement your solution here

	for i := 1; i < len(arr); i = i + 2 {
		if arr[i] > arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
		if (i + 1) <= len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

	}

	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[i]", arr[i])
	}
}

func indexArray(arr []int, size int) {
	//Implement your solution here
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[j] == i {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	for i := 0; i < size; i++ {
		if arr[i] != i {
			arr[i] = -1
		}
	}

	fmt.Println("Array: ", arr)
}

func SmallestPositiveMissingNumber(arr []int, size int) int {
	//Implement your solution here
	missing := -1
	exists := false
	for i := 0; i < size; i++ {
		exists = false
		for j := 0; j < size; j++ {
			if arr[j] == i+1 {
				exists = true
			}
		}
		if exists == false {
			missing = i + 1
			break
		}
	}

	return missing //Return -1 if missing number not found
}

func MaxMinArr(arr []int, size int) {
	//Implement your solution here
	min := 0
	max := 0
	for i := 0; i < size; i = i + 2 {
		for j := i; j < size; j++ {
			if arr[j] > arr[max] {
				max = j
			}
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
		arr[i+1], arr[min] = arr[min], arr[i+1]

		fmt.Println(arr[min], arr[max])
	}

	fmt.Println(arr)
}

func main() {
	// fmt.Println("N = 100, Number of instructions O(n) ::", foo(100))
	arr := []int{1, 2, 3, 4, 5, 6}

	// maxSum := []int{1, 2, 3, 4, 5, 10, 9, 8}
	// fmt.Println("Array Sum: ", sumArr(arr))
	// fmt.Println("===========================")
	// fmt.Println("Search Key: ", SequentialSearch(arr, 2))
	// fmt.Println("===========================")
	// fmt.Println("Max Subsequence: ", MaxSubArraySum(maxSum))
	// fmt.Println("===========================")
	// fmt.Println("Rotate: ", RotateArray(arr, 2))
	// fmt.Println("===========================")
	// fmt.Println("Wavy")
	// WaveArray(maxSum)
	// fmt.Println("===========================")
	// indexArray(arr, 10)
	// fmt.Println("===========================")
	// fmt.Println("Missing: ", SmallestPositiveMissingNumber(arr, 9))
	fmt.Println("===========================")
	MaxMinArr(arr, 6)
}
