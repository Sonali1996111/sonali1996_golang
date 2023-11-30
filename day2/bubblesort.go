//A bubble sorting algorithm commonly used to sort elements in a list or array.
//In this program by using for loop basic algorithm is
// arrenging a string of number or other elements in correct order .
//The bubble sort has complecity o(N).

package main

import "fmt"

func main() {
	abcd := []int{12, 33, 44, 21, 55, 88, 24, 43, 2, 7, 78, 90, 42}
	fmt.Println(Bubblesort(abcd))

}
func Bubblesort(abcd []int) []int {
	for i := 0; i < len(abcd)-1; i++ {
		for j := 0; j < len(abcd)-1; j++ {
			if abcd[j] > abcd[j+1] {
				abcd[j], abcd[j+1] = abcd[j+1], abcd[j]
			}
		}
	}
	return abcd
}
