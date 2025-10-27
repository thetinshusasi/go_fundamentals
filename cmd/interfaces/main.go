package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type shapename interface {
	name() string
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r circle) area() float64 {
	return 3.14 * math.Pow(r.radius, 2)
}

func (r circle) perimeter() float64 {
	return 2 * 3.14 * r.radius
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func display(s shape) {
	fmt.Println("this is the shape : ", s.area(), s.perimeter())
}

func displayUpdated(data interface{}) {

	switch v := data.(type) {
	case rect:
		fmt.Println("rect assertions : ", v.width, v.height)
		break
	case circle:
		fmt.Println("Circle assertions : ", v.radius)

	default:
		fmt.Println("No assertions ")
	}
}

func main() {

	rect1 := rect{
		width:  20,
		height: 4,
	}

	cir1 := circle{
		radius: 20,
	}

	var shape1 shape = rect{width: 10, height: 45}
	display(rect1)
	display(cir1)

	// type assertion

	instance, ok := shape1.(rect)
	if ok {
		fmt.Println("Type assertions : ", instance.width, instance.height)
	}

	displayUpdated(cir1)
	displayUpdated(rect1)

}
