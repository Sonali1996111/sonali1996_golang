//As a sequential search algorithm that starts at one end and goes through each elements of a list
//until the desired element is found,otherwise the search continues till the end of the data set.
//Iterator through a collection one element at a time.using for loop.
//complaxcity is o(N) for this perticular program.
//work for both sorted and unsorted array.

package main

import "fmt"

func printArray(array []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(array[i], "")
	}
	fmt.Println()
}
func linearSearch(array []int, size int, toFind int) int {
	for i := 0; i < size; i++ {
		if array[i] == toFind {
			return i
		}
	}
	return -1

}

func main() {
	array := []int{10, 5, 3, 7, 6, 12}
	var toSearch int
	toSearch = 3
	fmt.Println("golang program to find an element in an array using linear search")
	fmt.Println("array:")
	printArray(array, 6)
	index := linearSearch(array, 6, toSearch)

	if index == -1 {
		fmt.Println(toSearch, "is not present in the array")
	} else {
		fmt.Println(toSearch, "is present at index 2 in the array.")
	}
}
