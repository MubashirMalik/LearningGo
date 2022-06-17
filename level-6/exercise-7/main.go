package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("I am a first class citizen!")
	}
	f()
	g := f
	g()
}
