package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, "is", p.age)
}

func main() {
	x := person{
		first: "Pete",
		last:  "Bonker",
		age:   99,
	}

	x.speak()

}
