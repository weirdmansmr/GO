package main

import (
	"fmt"
)

func main() {
	totalPriceChecker()
}

func totalPriceChecker()  {
	price := 100
	fmt.Println("Price is", price, "dollars.")
	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	availableFunds := 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}

// func printResults() {
// 	fmt.Printf("Circle's square: %f sm^2 \n", calculateCircleArea(5))
// 	fmt.Printf("Rect square: %d \n", calculateRectangularSquare(5, 5))
// 	fmt.Printf("Rect permeter: %d \n", calculateRectangularPerimeter(5, 5))
// 	fmt.Printf("Triangle square: %f \n", calculateTriangleSquare(5, 3))
// }

// func calculateCircleArea(r int) float32 {
// 	return float32(r) * float32(r) * math.Pi
// }

// func calculateRectangularSquare(a int, b int) int {
// 	return a * b
// }

// func calculateRectangularPerimeter(a int, b int) int {
// 	return a*2 + b*2
// }

// func calculateTriangleSquare(b int, h int) float32 {
// 	return float32(h) * float32(b) / 2
// }