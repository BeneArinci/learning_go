package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func main() {
	t := triangle{}
	t.base = 5
	t.height = 6

	s := square{}
	s.side = 2

	printArea(t)

	printArea(s)

}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
