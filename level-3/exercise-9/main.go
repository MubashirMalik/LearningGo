package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("go to europe!")
	case "hockey":
		fmt.Println("go to australia!")
	case "cricket":
		fmt.Println("go to pakistan!")
	}
}
