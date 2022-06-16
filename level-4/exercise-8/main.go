package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, v1 := range m {
		fmt.Println("This is the record for: ", i)
		for j, v2 := range v1 {
			fmt.Println(j, v2)
		}
	}
}
