package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// %v is the value in the default format
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
