package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var target int = rand.IntN(100)
	var guess int
	fmt.Println("Guess a number between 1 and 100")
	fmt.Scan(&guess)
	i := 0
	for i = 1; guess != target; i++ {
		if guess < target {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")
		}
		fmt.Scan(&guess)
	}

	fmt.Println("Congratulations, you guessed the number after %d attempts", i)
}
