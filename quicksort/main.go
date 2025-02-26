// Package main quicksort
package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 8, 2, 5, 1, 9}
	n := len(arr)
	quicksort(arr, 0, n-1)
	fmt.Println(arr)
}
