package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello world!")
}

func foo() {
	defer func() {
		fmt.Println("foo defer ran")
	}()
	fmt.Println("foo ran")
}
