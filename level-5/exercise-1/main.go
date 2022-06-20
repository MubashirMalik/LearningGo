package main

import "fmt"

type person struct {
	first           string
	last            string
	icecreamFlavors []string
}

func main() {
	p1 := person{
		first: "Furqan",
		last:  "Tariq",
		icecreamFlavors: []string{
			"Strawberry", "Vanilla",
		},
	}
	p2 := person{
		first: "Mubashir",
		last:  "Ahmed",
		icecreamFlavors: []string{
			"Mango", "Kulfa",
		},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.icecreamFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.icecreamFlavors {
		fmt.Println(i, v)
	}
}
