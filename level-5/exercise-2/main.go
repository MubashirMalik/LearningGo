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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
