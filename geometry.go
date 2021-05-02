package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("perimeter: ", g.perimeter())
}

func main() {
	var shapeString string
	var shape geometry
	fmt.Println("Enter Shape ( 'circle' , 'rectangle' , 'square' ) : ")
	fmt.Scanln(&shapeString)
	switch shapeString {
	case "circle":
		fmt.Println("Enter Circle radius : ")
		var radius float64
		fmt.Scanln(&radius)
		shape = circle{radius: radius}

	case "square":
		fmt.Println("Enter Square side : ")
		var side float64
		fmt.Scanln(&side)
		shape = square{side: side}

	case "rectangle":
		fmt.Println("Enter rectangle height : ")
		var height float64
		fmt.Scanln(&height)
		fmt.Println("Enter rectangle width : ")
		var width float64
		fmt.Scanln(&width)
		shape = rectangle{height: height, width: width}
	default:
		fmt.Println("Invalid Shape entered : ")
	}

	measure(shape)
}
