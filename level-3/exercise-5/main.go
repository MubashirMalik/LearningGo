package main

import "fmt"

func main() {
	for i := 11; i < 100; i++ {
		fmt.Println(i % 4)
	}
}
