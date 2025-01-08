package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 5, Height: 10}
	triangle := Triangle{A: 3, B: 4, C: 5}

	PrintShapeDetails(circle)
	PrintShapeDetails(rectangle)
	PrintShapeDetails(triangle)
	PrintShapeDetails(10)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func PrintShapeDetails(s interface{}) {
	shape, ok := s.(Shape)
	if ok {
		fmt.Println("Area: ", shape.Area())
		fmt.Println("Perimeter: ", shape.Perimeter())

	} else {
		if isCircle(s) {
			fmt.Println("Circle doesn't implement Shape interface")
		} else if isRectangle(s) {
			fmt.Println("Rectangle doesn't implement Shape interface")
		} else if isTriangle(s) {
			fmt.Println("Triangle doesn't implement Shape interface")
		} else {
			fmt.Println("Unknown Shape")
			return
		}

	}

	DetectExactShape(s)

}

func DetectExactShape(s interface{}) {
	switch s.(type) {
	case Circle:
		fmt.Println("this is a cicle")
		c := s.(Circle)
		fmt.Println("Radius: ", c.Radius)
	case Triangle:
		fmt.Println("this is a triangle")
		t := s.(Triangle)
		fmt.Println(t.A, t.B, t.C)
	case Rectangle:
		fmt.Println("this is a rectangle")
		r := s.(Rectangle)
		fmt.Println(r.Height, r.Width)
	default:
		fmt.Println("Unknown shape")
	}
}

func isCircle(c interface{}) bool {
	//fmt.Println("testing circle")
	_, ok := c.(Circle)
	return ok
}

func isRectangle(r interface{}) bool {
	_, ok := r.(Rectangle)
	return ok
}

func isTriangle(t interface{}) bool {
	_, ok := t.(Triangle)
	return ok
}
