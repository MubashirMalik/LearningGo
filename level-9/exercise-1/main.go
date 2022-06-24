package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	fmt.Println("I am foo!")
	wg.Done()
}

func bar() {
	fmt.Println("I am bar!")
	wg.Done()
}
