package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	gs := 100
	wg.Add(gs)

	var mut sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()
			v := counter
			v++
			counter = v
			mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
