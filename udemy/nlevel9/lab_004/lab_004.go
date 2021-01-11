package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 0
	g := 10
	wg.Add(g)
	var mut sync.Mutex

	for i := 0; i < g; i++ {
		go func() {
			mut.Lock()
			fmt.Println("x:", x)
			v := x
			v++
			x = v
			mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
