package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var repeat bool = true
	for repeat {
		var rang, choice, nb_tries int = 0, 0, 0

		for rang == 0 {
			fmt.Println("Please choose from the menu a level of difficulty and enter the number: \n 1.Easy \n 2.Medium \n 3. Hard")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				rang = 50
				fmt.Println("You chose the easy level, the number is between 0 and 50")
			case 2:
				rang = 100
				fmt.Println("You chose the medium level, the number is between 0 and 100")
			case 3:
				rang = 200
				fmt.Println("You chose the hard level, the number is between 0 and 200")
			default:
				fmt.Println("Invalid Choice")
			}
		}
		fmt.Println("In how many attempts you think you'll guess the correct number")
		for nb_tries <= 0 {
			fmt.Print("Enter the max attempts : ")
			fmt.Scan(&nb_tries)
			switch {
			case nb_tries > 0:
				fmt.Printf("We'll see if you can guess it in %d attempts :) \n", nb_tries)
			default:
				fmt.Println("Invalid choice")
			}
		}

		var target int = rand.IntN(rang)
		var guess int
		var score int = rang
		fmt.Println("You can start guessing the number")
		fmt.Scan(&guess)
		i := 0
		for i = 1; guess != target && i < nb_tries; i++ {
			if guess > rang || guess < 0 {
				fmt.Printf("The number is between 0 and %d \n", rang)
			} else if guess < target {
				score = target - guess
				fmt.Println("Too low")
			} else {
				score = guess - target
				fmt.Println("Too high")
			}
			fmt.Scan(&guess)
		}
		if guess == target {
			fmt.Printf("Congratulations *\\0/*, you guessed the number after %d attempts \n", i)
		} else {
			fmt.Printf("You didn't guess the number, the correct number was %d \n", target)
			fmt.Printf("You were %d away from the correct number \n", score)
		}

		fmt.Println("It was fun, don't you think , do you want to play again ? type 'yes' if so ")
		var answer string
		fmt.Scan(&answer)
		if answer != "yes" {
			fmt.Println("Oh it wasn't that fun, see you next time")
			repeat = false
		}

	}
}
