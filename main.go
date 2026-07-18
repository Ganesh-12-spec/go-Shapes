package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Solid interface {
	Volume() float64
}
type ShapeRequest struct{
	Type string `json:"type"`
	Radius float64 `json:"radius"`
	Width float64  `json:"width"`
	Length float64 `json:"length"`
	Base   float64 `json:"base"`
	Height float64 `json:"height"`
	Side1  float64 `json:"side1"`
	Side2  float64 `json:"side2"`
	Side3  float64 `json:"side3"`

}




type ShapeResponse struct {
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}

type Sphere struct{
	radius float64
}

type Circle struct{
	radius float64
}
type Cylinder struct {
	radius float64
	height float64
}

func (cy Cylinder) Volume() float64 {
	return math.Pi * cy.radius * cy.radius * cy.height
}

type Cone struct {
	radius float64
	height float64
}

func (co Cone) Volume() float64 {
	return (1.0 / 3.0) * math.Pi * co.radius * co.radius * co.height
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

func (s Sphere) Volume() float64{
	return (4.0/3.0)* math.Pi * s.radius *  s.radius * s.radius
}
func TotalArea(shapes []Shape) float64{
	count := 0.0
	for _, shape := range shapes{
		count += shape.Area()
	}
	return count
}
func LargestArea(shapes []Shape) Shape{
    largest := shapes[0]
		for _, shape := range shapes{
			if shape.Area() > largest.Area(){
				largest = shape
			}
		}
		return largest

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
		total := TotalArea(shapes)
		fmt.Println("Total:", total)

		largest := LargestArea(shapes)
    fmt.Println("Largest shape area:", largest.Area())

		sph := Sphere{radius: 3}
    fmt.Println("Sphere Volume:", sph.Volume())

		cyl := Cylinder{radius: 3, height: 5}
    cone := Cone{radius: 3, height: 5}

    fmt.Println("Cylinder Volume:", cyl.Volume())
    fmt.Println("Cone Volume:", cone.Volume())

		http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
			var req ShapeRequest
			json.NewDecoder(r.Body).Decode(&req)

			var shape Shape 

			if req.Type == "circle"{
				shape = Circle{radius: req.Radius}
			}else if req.Type == "rectangle"{
				shape = Rectangle{width: req.Width , length: req.Length}
			}else if req.Type == "triangle"{
				shape = Triangle{height: req.Height,base: req.Base,side1: req.Side1,side2:req.Side2, side3: req.Side3}
			}else{
				fmt.Print("not valid")
			}

			response := ShapeResponse{
				Area:     shape.Area(),
				Perimeter: shape.Perimeter(),
			}
			json.NewEncoder(w).Encode(response)
		})

		http.ListenAndServe(":8080",nil)
}