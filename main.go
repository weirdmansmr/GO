package main

import (
	"fmt"
	"math"
)

func main() {
	printResults()
}

func printResults() {
	fmt.Printf("Circle's square: %f sm^2 \n", calculateCircleArea(5))
	fmt.Printf("Rect square: %d \n", calculateRectangularSquare(5, 5))
	fmt.Printf("Rect permeter: %d \n", calculateRectangularPerimeter(5, 5))
	fmt.Printf("Triangle square: %f \n", calculateTriangleSquare(5, 3))
}

func calculateCircleArea(r int) float32 {
	return float32(r) * float32(r) * math.Pi
}

func calculateRectangularSquare(a int, b int) int {
	return a * b
}

func calculateRectangularPerimeter(a int, b int) int {
	return a*2 + b*2
}

func calculateTriangleSquare(b int, h int) float32 {
	return float32(h) * float32(b) / 2
}