package main

import "fmt"

func main() {

	var a, b, c = 11, 33, 12

	if a > b {
		if a > c {

			fmt.Println("a is greater")

		} else {
			fmt.Println("c is greater")

		}
	} else if b > c {
		fmt.Println("b is greater")

	} else {
		fmt.Println("c is greater")
	}
	fmt.Println("hello", "playground")
}
