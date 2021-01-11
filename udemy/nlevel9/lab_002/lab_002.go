package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Name is", p.name)
}

type human interface {
	speak()
}

func main() {
	p1 := &person{"Tom"}
	//p2 := person{"Jerry"}

	saysomething(p1)
	//saysomething(p2)
}

func saysomething(h human) {
	h.speak()
}
