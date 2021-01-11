package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("Func")
	}
	f()

	x := func(s string) {
		fmt.Println(s)
	}

	x("HELLO")
}
