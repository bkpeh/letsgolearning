package main

import "fmt"

type SQUARE struct {
	length int
	width  int
}

type CIRCLE struct {
	radius int
}

func (sq SQUARE) area() float64 {
	return float64(sq.length) * float64(sq.width)
}

func (ci CIRCLE) area() float64 {
	return 3.142 * float64(ci.radius) * 2.0
}

type SHAPE interface {
	area() float64
}

func main() {
	sq := SQUARE{2, 3}
	ci := CIRCLE{3}

	INFO(sq)
	INFO(ci)
}

func INFO(s SHAPE) {
	fmt.Println(s.area())
}
