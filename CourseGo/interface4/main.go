package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func main() {
	rect := Rectangle{width: 5, height: 4}
	// circle := Circle{radius: 4}

	describeShape(rect)

}

func describeShape(g Geometry) {
	fmt.Println("Area", g.Area())
	fmt.Println("Perimeter", g.Perimeter())
}

type CalculationError struct {
	message string
}

func (ce CalculationError) Error() string {
	return ce.message
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{message: "Invalid input"}
	}
	return math.Sqrt(val), nil
}
