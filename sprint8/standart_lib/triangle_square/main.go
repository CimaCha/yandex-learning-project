package main

import (
	"fmt"
	"math"
)

func main() {
	// Площадь прямоугольного треугольника с катетами 3 и 4
	fmt.Println(triangleArea(3.0, 4.0, 90.0))
}

func triangleArea(a, b, angleDegrees float64) float64 {
	// Вычислите площадь треугольника
	angleRadians := angleDegrees * math.Pi / 180
	s := a * b / 2 * math.Sin(angleRadians)
	return s
}
