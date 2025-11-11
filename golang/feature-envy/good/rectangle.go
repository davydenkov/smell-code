package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r Rectangle) GetWidth() float64 {
	return r.width
}

func (r Rectangle) GetHeight() float64 {
	return r.height
}

func (r Rectangle) CalculateArea() float64 {
	return r.width * r.height
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) IsSquare() bool {
	return r.width == r.height
}

func (r Rectangle) CalculateDiagonal() float64 {
	return math.Sqrt(r.width*r.width + r.height*r.height)
}

func (r Rectangle) GetAspectRatio() float64 {
	return r.width / r.height
}

type GeometryUtils struct{}

// Utility methods that don't belong to Rectangle can stay here
func (gu GeometryUtils) CalculateDistanceBetweenPoints(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func (gu GeometryUtils) CalculateAngle(opposite, adjacent float64) float64 {
	return math.Atan(opposite / adjacent)
}

func main() {
	rect := NewRectangle(10, 5)

	area := rect.CalculateArea()
	perimeter := rect.CalculatePerimeter()
	isSquare := rect.IsSquare()
	diagonal := rect.CalculateDiagonal()
	aspectRatio := rect.GetAspectRatio()

	println("Area:", area)
	println("Perimeter:", perimeter)
	println("Is Square:", isSquare)
	println("Diagonal:", diagonal)
	println("Aspect Ratio:", aspectRatio)

	utils := GeometryUtils{}
	distance := utils.CalculateDistanceBetweenPoints(0, 0, 3, 4)
	angle := utils.CalculateAngle(3, 4)

	println("Distance between points:", distance)
	println("Angle:", angle)
}
