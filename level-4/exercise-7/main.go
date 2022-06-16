package main

import "fmt"

func main() {

	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	s3 := [][]string{s1, s2}
	for _, v1 := range s3 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
