package main

import "fmt"

func main() {
	x := 1
	func() {
		x := 5
		fmt.Println(x)
	}()
	fmt.Println(x)
}
