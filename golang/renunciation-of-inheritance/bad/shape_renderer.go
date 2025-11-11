package main

import (
	"fmt"
	"math"
)

type CircleRenderer struct {
	color string
}

func NewCircleRenderer(color string) *CircleRenderer {
	if color == "" {
		color = "black"
	}
	return &CircleRenderer{color: color}
}

func (cr CircleRenderer) Render(radius float64) {
	fmt.Printf("Rendering circle with radius %f in %s\n", radius, cr.color)
}

func (cr CircleRenderer) GetArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func (cr CircleRenderer) GetColor() string {
	return cr.color
}

func (cr *CircleRenderer) SetColor(color string) {
	cr.color = color
}

type RectangleRenderer struct {
	color string
}

func NewRectangleRenderer(color string) *RectangleRenderer {
	if color == "" {
		color = "black"
	}
	return &RectangleRenderer{color: color}
}

func (rr RectangleRenderer) Render(width, height float64) {
	fmt.Printf("Rendering rectangle %fx%f in %s\n", width, height, rr.color)
}

func (rr RectangleRenderer) GetArea(width, height float64) float64 {
	return width * height
}

func (rr RectangleRenderer) GetColor() string {
	return rr.color
}

func (rr *RectangleRenderer) SetColor(color string) {
	rr.color = color
}

type TriangleRenderer struct {
	color string
}

func NewTriangleRenderer(color string) *TriangleRenderer {
	if color == "" {
		color = "black"
	}
	return &TriangleRenderer{color: color}
}

func (tr TriangleRenderer) Render(base, height float64) {
	fmt.Printf("Rendering triangle with base %f and height %f in %s\n", base, height, tr.color)
}

func (tr TriangleRenderer) GetArea(base, height float64) float64 {
	return 0.5 * base * height
}

func (tr TriangleRenderer) GetColor() string {
	return tr.color
}

func (tr *TriangleRenderer) SetColor(color string) {
	tr.color = color
}

// Code duplication in color management and structure - should use composition/interfaces!

func main() {
	circle := NewCircleRenderer("red")
	circle.Render(5.0)
	fmt.Printf("Circle area: %.2f\n", circle.GetArea(5.0))

	rectangle := NewRectangleRenderer("blue")
	rectangle.Render(10.0, 8.0)
	fmt.Printf("Rectangle area: %.2f\n", rectangle.GetArea(10.0, 8.0))

	triangle := NewTriangleRenderer("green")
	triangle.Render(6.0, 4.0)
	fmt.Printf("Triangle area: %.2f\n", triangle.GetArea(6.0, 4.0))
}
