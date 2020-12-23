package main

import "fmt"

func main() {

	fmt.Println("Start")
	defer f2()

	f1()

	fmt.Println("End")
}

func f1() {
	fmt.Println("F1")
}

func f2() {
	fmt.Println("Defer")
}
