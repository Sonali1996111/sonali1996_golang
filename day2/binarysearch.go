//it divides the the array into two halves at every iteration to narrow down the search space .
//A search algorithm that finds the position of a target value within a sorted array.by using for,and if else statment.
//The time complexcity is log2n.
//work only for sorted arrays.

package main

import "fmt"

func Binarysearch(key int, stock []int) bool {
	low := 0
	high := len(stock) - 1
	for low <= high {
		medium := (low + high) / 2

		if stock[medium] < key {
			low = medium + 1
		} else {
			high = medium - 1
		}
	}
	if low == len(stock) || stock[low] != key {
		return false
	}
	return true
}
func main() {
	items := []int{1, 2, 7, 9, 8, 45, 32, 99, 66, 3}
	fmt.Println(Binarysearch(44, items))
}
