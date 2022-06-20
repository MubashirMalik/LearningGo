package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Mubashir",
		friends: map[string]int{
			"Furqan": 420,
			"Rehan":  69,
			"Haris":  0,
		},
		favDrinks: []string{
			"Mango Shake", "Tea", "Coke",
		},
	}
	fmt.Println(s.first)
}
