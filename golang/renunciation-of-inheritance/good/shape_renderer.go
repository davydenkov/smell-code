package main

import (
	"fmt"
	"math"
)

// ShapeRenderer interface defines the common contract
type ShapeRenderer interface {
	Render()
	GetArea() float64
	GetColor() string
	SetColor(color string)
}

// BaseShape provides common functionality
type BaseShape struct {
	color string
}

func NewBaseShape(color string) BaseShape {
	if color == "" {
		color = "black"
	}
	return BaseShape{color: color}
}

func (bs BaseShape) GetColor() string {
	return bs.color
}

func (bs *BaseShape) SetColor(color string) {
	bs.color = color
}

func (bs BaseShape) getRenderPrefix() string {
	return fmt.Sprintf("Rendering in %s", bs.color)
}

// CircleRenderer implements ShapeRenderer
type CircleRenderer struct {
	BaseShape
	radius float64
}

func NewCircleRenderer(radius float64, color string) *CircleRenderer {
	return &CircleRenderer{
		BaseShape: NewBaseShape(color),
		radius:    radius,
	}
}

func (cr CircleRenderer) Render() {
	fmt.Printf("%s circle with radius %.2f\n", cr.getRenderPrefix(), cr.radius)
}

func (cr CircleRenderer) GetArea() float64 {
	return math.Pi * cr.radius * cr.radius
}

func (cr CircleRenderer) GetRadius() float64 {
	return cr.radius
}

// RectangleRenderer implements ShapeRenderer
type RectangleRenderer struct {
	BaseShape
	width  float64
	height float64
}

func NewRectangleRenderer(width, height float64, color string) *RectangleRenderer {
	return &RectangleRenderer{
		BaseShape: NewBaseShape(color),
		width:     width,
		height:    height,
	}
}

func (rr RectangleRenderer) Render() {
	fmt.Printf("%s rectangle %.2fx%.2f\n", rr.getRenderPrefix(), rr.width, rr.height)
}

func (rr RectangleRenderer) GetArea() float64 {
	return rr.width * rr.height
}

func (rr RectangleRenderer) GetWidth() float64 {
	return rr.width
}

func (rr RectangleRenderer) GetHeight() float64 {
	return rr.height
}

// TriangleRenderer implements ShapeRenderer
type TriangleRenderer struct {
	BaseShape
	base   float64
	height float64
}

func NewTriangleRenderer(base, height float64, color string) *TriangleRenderer {
	return &TriangleRenderer{
		BaseShape: NewBaseShape(color),
		base:      base,
		height:    height,
	}
}

func (tr TriangleRenderer) Render() {
	fmt.Printf("%s triangle with base %.2f and height %.2f\n", tr.getRenderPrefix(), tr.base, tr.height)
}

func (tr TriangleRenderer) GetArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (tr TriangleRenderer) GetBase() float64 {
	return tr.base
}

func (tr TriangleRenderer) GetHeight() float64 {
	return tr.height
}

// ShapeRendererManager demonstrates polymorphism
type ShapeRendererManager struct{}

func (srm ShapeRendererManager) RenderAllShapes(shapes []ShapeRenderer) {
	for _, shape := range shapes {
		shape.Render()
		fmt.Printf("Area: %.2f\n", shape.GetArea())
	}
}

func main() {
	// Create different shapes using composition and interfaces
	circle := NewCircleRenderer(5.0, "red")
	rectangle := NewRectangleRenderer(10.0, 8.0, "blue")
	triangle := NewTriangleRenderer(6.0, 4.0, "green")

	// Store in slice of interface type for polymorphism
	shapes := []ShapeRenderer{circle, rectangle, triangle}

	// Demonstrate polymorphism
	manager := ShapeRendererManager{}
	manager.RenderAllShapes(shapes)

	// Change colors using common interface
	circle.SetColor("purple")
	fmt.Printf("Circle is now %s\n", circle.GetColor())

	fmt.Println("\nGo uses interfaces and composition instead of inheritance:")
	fmt.Println("- ShapeRenderer interface defines common contract")
	fmt.Println("- BaseShape provides shared functionality through embedding")
	fmt.Println("- Each shape implements the interface with its specific behavior")
	fmt.Println("- Polymorphism achieved through interface types")
}
