package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, _ := range x {
		for ii, _ := range x[i] {
			fmt.Print(x[i][ii], " ")
		}
		fmt.Println()
	}
}
