package main

import "fmt"

func main() {
	x := 1980

	for {
		if x > 2020 {
			break
		}

		fmt.Println(x)
		x++
	}
}
