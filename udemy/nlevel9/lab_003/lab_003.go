package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 0
	g := 10
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			fmt.Println("x:", x)
			v := x
			runtime.Gosched()
			v++
			x = v
			wg.Done()
		}()
	}
	wg.Wait()
}
