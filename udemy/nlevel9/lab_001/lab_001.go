package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

//var wg2 sync.WaitGroup

func main() {
	fmt.Println("Main")

	go foo()
	go bar()
	print()
	wg1.Add(2)
	wg1.Wait()

}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo:", i)
	}
	wg1.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("Bar:", i)
	}
	wg1.Done()
}

func print() {
	for i := 0; i < 20; i++ {
		fmt.Println("print:", i)
	}
}
