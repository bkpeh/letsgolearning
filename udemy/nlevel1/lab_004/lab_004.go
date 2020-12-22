package main

import "fmt"

type hey int

func main() {
	var x hey

	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 42
	fmt.Println(x)

}
