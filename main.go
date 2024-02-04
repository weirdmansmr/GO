package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
func main() {
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
		break
		}
}
if !success {
	fmt.Println("Sorry, you didn't guess my number. It was:", target)
}
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