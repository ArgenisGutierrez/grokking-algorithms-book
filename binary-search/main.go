package main

import (
	"errors"
	"fmt"
)

func main() {
	my_list := []int{0, 2, 4, 6, 8}
	num, err := binarySearch(1, my_list)
	if err != nil {
		fmt.Println("None")
		return
	}

	fmt.Println(num)

}

func binarySearch(item int, list []int) (int, error) {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high)
		guess := list[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, errors.New("None")
}
