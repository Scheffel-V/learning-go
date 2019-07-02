package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	square1 := square{10}
	triangle1 := triangle{10, 10}

	printArea(square1)
	printArea(triangle1)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.height * t.base) / 2
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

// 1. No errors
