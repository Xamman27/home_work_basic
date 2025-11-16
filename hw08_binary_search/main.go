package main

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		middle := start + (end-start)/2
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

func main() {
	numbers := make([]int, 1001)
	for i := 0; i <= 1000; i++ {
		numbers[i] = i
	}
	target := 760
	my_found := BinarySearch(numbers, target)
	println(my_found)
}
