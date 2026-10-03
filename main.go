package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

var _ Shape = (*Circle)(nil)

type Square struct {
	Side float64
}

var _ Shape = (*Square)(nil)

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

// type Shape interface { ... }
// type Circle struct { ... }
// func (c Circle) Area() float64 { ... }

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	kind := sc.Text()
	sc.Scan()
	dim, _ := strconv.ParseFloat(sc.Text(), 64)
	var s Shape
	_ = kind
	_ = dim
	switch kind {
	case "circle":
		s = &Circle{Radius: dim}
	case "square":
		s = &Square{Side: dim}

	}
	// s = ... based on kind
	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
