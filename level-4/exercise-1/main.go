package main

import "fmt"

func main() {
	arr := [5]int{31, 32, 33, 34, 35}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)
}
