package main

import "fmt"

func main() {

	var sum float64
	var nb_grades int

	fmt.Println("Enter grades, -1 to quit")

	for {
		var grade float64
		fmt.Println("Enter a grade:")
		fmt.Scan(&grade)
		if grade == -1 {
			break
		}
		sum += grade
		nb_grades++
	}

	var res float64 = sum / float64(nb_grades)

	fmt.Println("The average grade is: ", res)

	fmt.Print("The grade letter is: ")
	switch {
	case res >= 90:
		fmt.Println("A")
	case res >= 80:
		fmt.Println("B")
	case res >= 70:
		fmt.Println("C")
	case res >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

}
