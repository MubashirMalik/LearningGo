package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (s square) area() float64 {
	return (s.length * s.width)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 1.67,
	}
	s := square{
		length: 2.31,
		width:  1.76,
	}
	fmt.Print("Area of circle: ")
	info(c)
	fmt.Print("Area of square: ")
	info(s)
}
