package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var x int64 = 0
	g := 10

	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			atomic.AddInt64(&x, 1)
			//runtime.Gosched()
			fmt.Println("x:", atomic.LoadInt64(&x))
			wg.Done()
		}()
	}
	wg.Wait()
}
