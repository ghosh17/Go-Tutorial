//Interfaces in Go
package main

import "fmt"

func main() {
	cir := Circle{radius: 5}
	sq := Square{length: 4}
	a := getArea(cir)

	p := getPerimeter(sq)
	fmt.Println(a)
	fmt.Print("\n")
	fmt.Println(p)
}

//Declare the Interfaces

type Shape interface {
	area() float64
	perimeter() float64
}

//
type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

//
func getArea(shape Shape) float64 {
	return shape.area()
}

func getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}

//

func (circle Circle) perimeter() float64 {
	return 2 * 3.14 * circle.radius
}

func (sq Square) perimeter() float64 {
	return 4 * sq.length
}

func (circle Circle) area() float64 {
	return 3.14 * circle.radius * circle.radius
}

func (sq Square) area() float64 {
	return sq.length * sq.length
}
