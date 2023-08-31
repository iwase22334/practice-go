package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Square struct {
	w float64
	h float64
}

func (c Circle) GetArea() float64 {
	return c.r * c.r * math.Pi
}

func (s Square) GetArea() float64 {
	return s.w * s.h
}

type Figure interface {
	GetArea() float64
}

func PrintArea(f Figure) {
	fmt.Printf("Area: %v\n", f.GetArea())
}

func main() {
	c := Circle{r: 1.0}
	s := Square{w: 1.0, h: 2.0}

	PrintArea(c)
	PrintArea(s)
}
