package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println()
	// select a difficulty level
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 attempts)")
	fmt.Println("2. Medium (5 attempts)")
	fmt.Println("3. Hard (3 attempts)")
	fmt.Println()
	fmt.Print("Enter your choice: ")

	randomNumber := rand.IntN(100) + 1
	if randomNumber > 100 {
		randomNumber = 100
	}

	var level string
	var chances int

	for {
		if !scanner.Scan() {
			return
		}

		fmt.Println()

		choice := scanner.Text()

		switch choice {
		case "1":
			level = "Easy"
			chances = 10
		case "2":
			level = "Medium"
			chances = 5
		case "3":
			level = "Hard"
			chances = 3
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
			continue
		}
		break
	}

	fmt.Println("Great! You have selected the", level, "diggiculty level.")
	fmt.Println("Let's start the game!")
	fmt.Println()

	// game loop
	for count := 1; count <= chances; count++ {

		fmt.Print("Enter your guess: ")

		if !scanner.Scan() {
			break
		}

		guess := scanner.Text()

		gs, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			fmt.Println()
			count--
			continue
		}

		if gs == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", count)
			fmt.Println()
			break
		} else if gs < randomNumber {
			fmt.Printf("Incorrect. The number is greater than %s\n", guess)
			fmt.Println()
		} else if gs > randomNumber {
			fmt.Printf("Incorrect. The number is less than %s\n", guess)
			fmt.Println()
		}
	}

	fmt.Println("Sorry, you've used all your chances. The correct number was", randomNumber)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
