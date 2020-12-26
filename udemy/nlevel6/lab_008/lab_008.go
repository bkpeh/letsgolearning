package main

import "fmt"

func main() {
	f := ff()

	fmt.Println(f())
}

func ff() func() string {
	return func() string {
		return "FF Function"
	}
}
