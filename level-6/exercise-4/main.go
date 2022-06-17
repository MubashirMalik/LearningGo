package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("I am %v %v and I am %v years old.", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Mubashir",
		last:  "Ahmed",
		age:   23,
	}
	p.speak()
}
