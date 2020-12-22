package main

import "fmt"

func main() {
	x := 100

	fmt.Println(x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)

	x = x << 1

	fmt.Println(x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
}
