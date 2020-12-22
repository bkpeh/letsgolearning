package main

import "fmt"

type faveSport string

func main() {
	var s faveSport = "Football"

	switch s {
	case "Tennis":
		fmt.Println("Tennis")
	case "Football":
		fmt.Println("Football")
	default:
		fmt.Println("None")
	}
}
