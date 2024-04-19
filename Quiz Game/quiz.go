package main

import (
	"fmt"
)

func main() {
	questions := 3
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name, "!")
	fmt.Println("How old are you?")
	var age uint
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("Your are to young to play this game. Goodbye!")
		return
	} else if age > 150 {
		fmt.Println("You are too old to play this game. Goodbye!")
		return
	} else {
		fmt.Println("You are old enough to play this game. Let's start!")
	}
	var score int
	fmt.Println("Your score is ", score)

	fmt.Println("What is the capital of France?")
	var answer1 string
	fmt.Scanln(&answer1)
	if answer1 == "Paris" || answer1 == "paris" {
		score++
		fmt.Println("Correct! Your score is ", score)
	} else {
		fmt.Println("Incorrect! Your score is ", score)
	}

	fmt.Println("What is 12 - 5?")
	var answer2 int
	fmt.Scan(&answer2)
	if answer2 == 7 {
		score++
		fmt.Println("Correct! Your score is ", score)
	} else {
		fmt.Println("Incorrect! Your score is ", score)
	}

	fmt.Println("What is 5 * 5?")
	var answer3 int
	fmt.Scan(&answer3)
	if answer3 == 25 {
		score++
		fmt.Println("Correct! Your score is ", score)
	} else {
		fmt.Println("Incorrect! Your score is ", score)
	}

	fmt.Printf("You scored %d out of 1\n", score)
	percent := (float64(score) / float64(questions)) * 100
	fmt.Printf("You scored %v%%", percent)
}
