package main

import "fmt"

type shape interface {
	getArea() float64
	getPerimeter() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (t triangle) getPerimeter() float64 {
	return 2*t.height + t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (s square) getPerimeter() float64 {
	return 4 * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func printPerimeter(sh shape) {
	fmt.Println(sh.getPerimeter())
}

func main() {
	t := triangle{base: 10, height: 15}
	s := square{sideLength: 35}
	printArea(t)
	printPerimeter(t)
	printArea(s)
	printPerimeter(s)
}
