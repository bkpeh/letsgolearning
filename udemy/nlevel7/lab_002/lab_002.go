package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) {
	p.first = "Changed"
}

func main() {
	p := person{
		"Person",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
