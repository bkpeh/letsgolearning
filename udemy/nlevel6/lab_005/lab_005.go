package main

import "fmt"

type SQUARE struct {
	length int
	width  int
}

type CIRCLE struct {
	radius int
}

func (sq SQUARE) area() int {
	return sq.length * sq.width
}

func (ci CIRCLE) area() float64 {
	return 3.142 * ci.radius * 2
}

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
