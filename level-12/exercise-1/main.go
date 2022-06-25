package main

import (
	"fmt"

	"github.com/MubashirMalik/LearningGo/level-12/exercise-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
