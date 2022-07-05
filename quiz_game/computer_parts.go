/*
Quiz game developed starting from "Tech With Tim" YouTube video:
https://www.youtube.com/watch?v=LHhsNa_Kgns&list=PL0JfnK-pSXSCAjDRAA5lIsor3ly3-imy1&index=2
This code is part of my daily practice to learn Go as a programming language.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")

	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	num_questions := 0

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if strings.ToUpper(answer)+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!, The best graphic card is RTX 3090")
	}
	num_questions++

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Printf("Incorrect!, The number of cores is 12, and you typed %v.\n", cores)
	}
	num_questions++

	fmt.Printf("You scored %v out of %v questions.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
