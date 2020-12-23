package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	x := person{
		first:   "Pete",
		last:    "Bonker",
		flavors: []string{"Van", "Mint", "App"},
	}

	y := person{
		first:   "Jess",
		last:    "Walker",
		flavors: []string{"Choc", "Stra", "Lime"},
	}

	fmt.Println(x.first, x.last)
	for _, v := range x.flavors {
		fmt.Println(v)
	}

	fmt.Println(y.first, y.last)
	for _, v := range y.flavors {
		fmt.Println(v)
	}
}
