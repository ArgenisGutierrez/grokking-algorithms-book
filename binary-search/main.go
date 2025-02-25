package main

import "fmt"

func main() {
	my_list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	num := binarySearch(my_list, 3)
	res, count := binarySearchRecursive(my_list, 15)
	fmt.Println(num)
	fmt.Println(res, count)
}

func binarySearch(list []int, num int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		switch {
		case guess == num:
			return mid
		case guess > num:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return -1
}

func binarySearchRecursive(list []int, num int) (result, count int) {
	mid := len(list) / 2
	switch {
	case len(list) == 0:
		result = -1
	case list[mid] > num:
		result, count = binarySearchRecursive(list[:mid], num)
	case list[mid] < num:
		result, count = binarySearchRecursive(list[mid:], num)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}
	count++
	return
}
