package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(in ...int) int {
	var total int

	for _, v := range in {
		total += v
	}
	return total
}

func bar(in []int) int {
	var total int

	for _, v := range in {
		total += v
	}
	return total
}
