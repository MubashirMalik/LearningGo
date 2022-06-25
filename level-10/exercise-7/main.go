package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := puts(q)
	pulls(c, q)
	fmt.Println("exiting main")
}

func puts(q chan<- int) <-chan int {
	c := make(chan int)
	for gs := 0; gs < 10; gs++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			q <- 1
		}()
	}
	return c
}

func pulls(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("pulled number:", v)
		case <-q:
			return
		}
	}
}
