package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

type GeometryUtils struct{}

// This method has feature envy - it accesses many fields of Rectangle
func (gu GeometryUtils) CalculateArea(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (gu GeometryUtils) CalculatePerimeter(rectangle *Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (gu GeometryUtils) IsSquare(rectangle *Rectangle) bool {
	return rectangle.Width == rectangle.Height
}

func (gu GeometryUtils) CalculateDiagonal(rectangle *Rectangle) float64 {
	return math.Sqrt(rectangle.Width*rectangle.Width + rectangle.Height*rectangle.Height)
}

func (gu GeometryUtils) GetAspectRatio(rectangle *Rectangle) float64 {
	return rectangle.Width / rectangle.Height
}

func main() {
	utils := GeometryUtils{}
	rect := NewRectangle(10, 5)

	area := utils.CalculateArea(rect)
	perimeter := utils.CalculatePerimeter(rect)
	isSquare := utils.IsSquare(rect)
	diagonal := utils.CalculateDiagonal(rect)
	aspectRatio := utils.GetAspectRatio(rect)

	println("Area:", area)
	println("Perimeter:", perimeter)
	println("Is Square:", isSquare)
	println("Diagonal:", diagonal)
	println("Aspect Ratio:", aspectRatio)
}
