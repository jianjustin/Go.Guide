package example

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	weight, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.weight * r.height
}

func (r rect) perim() float64 {
	return 2*r.weight + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
