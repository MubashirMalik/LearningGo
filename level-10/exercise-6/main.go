package main

import "fmt"

func main() {
	q := make(chan int)
	c := puts(q)
	pulls(c, q)
	fmt.Println("exiting main")
}

func puts(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
	}()
	return c
}

func pulls(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
