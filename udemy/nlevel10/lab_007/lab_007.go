package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for ii := 0; ii < 10; ii++ {
				c <- ii
			}
		}()
	}

	for ii := 1; ii <= 100; ii++ {
		fmt.Println(ii, <-c)
	}
}
