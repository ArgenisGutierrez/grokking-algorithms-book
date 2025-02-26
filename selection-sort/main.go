// Package main Selection sort algorithm
package main

import "fmt"

func selectionSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}

	return slice
}

func main() {
	myList := []int{5, 1, 6, 3, 7, 9}
	newList := selectionSort(myList)
	fmt.Println(newList)

}
