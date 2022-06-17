package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	s1 := sum(x...)
	fmt.Println(s1)
	s2 := even(sum, x...)
	fmt.Println(s2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(y ...int) int, x ...int) int {
	var xi []int
	for _, v := range x {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	return f(xi...)
}
