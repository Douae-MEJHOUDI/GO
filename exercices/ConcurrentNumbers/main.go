package main

import (
	"fmt"
	"sync"
)

func Square(x int) {
	fmt.Printf("Square of %d is %d\n", x, x*x)
}

func main() {

	var wg sync.WaitGroup

	listNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range listNumbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Square(v)
		}()

	}
	wg.Wait()

}
