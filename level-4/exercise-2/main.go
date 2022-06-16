package main

import "fmt"

func main() {
	arr := []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)
}
