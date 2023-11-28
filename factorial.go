package main

import "fmt"

//7*6*5*4*3*2*1
func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result

}

func factr(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factr(n-1)
	}
}

func main() {
	fmt.Println(fact(7))

}
