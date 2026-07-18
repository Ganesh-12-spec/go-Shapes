package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{
	radius float64
}
type Rectangle struct{
	width  float64
	length float64
}
type Triangle struct {
	base   float64
	height float64
	side1  float64
	side2  float64
	side3  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r Rectangle) Area() float64{
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main(){
    c := Circle{radius: 5}
    r := Rectangle{width: 6, length: 5}
    t := Triangle{base: 4, height: 3, side1: 3, side2: 4, side3: 5}
   

		shapes := []Shape{
			c, r, t,
		}
		for _, shape := range shapes {
			fmt.Println("Area:", shape.Area())
			fmt.Println("Perimeter:", shape.Perimeter())
		}
}