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

	per := map[string]person{
		"Pete": x,
		"Jess": y,
	}

	for _, v := range per {
		fmt.Println(v.first, v.last)
		for _, vv := range v.flavors {
			fmt.Println(vv)
		}
	}
}
