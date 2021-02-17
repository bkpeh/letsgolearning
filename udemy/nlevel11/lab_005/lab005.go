package main

import (
	"fmt"
)

func main() {
	foo("From Main")
	fmt.Println("Ending")
}

func foo(s string) {
	fmt.Println("This is foo.", s)
}
