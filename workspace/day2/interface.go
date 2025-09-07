package main

import (
	"fmt"
	"math"
)

// named interface. In other words goemetry is the interface type
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/* define required methods to implement geometry interface */
func (r rect) area() float64 {
	return r.width * r. height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/**/

// g holds interface values, which is either rect or circle
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	// assertion
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius",c.radius)
	} else {
		fmt.Println("I am else")
	}
}

func main() {
	// instance of struct
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("measure(r)")
	measure(r)
	fmt.Println("measure(c)")
	measure(c)
	fmt.Println("detectCircle(r)")
	detectCircle(r)
	fmt.Println("detectCircle(c)")
	detectCircle(c)
}
