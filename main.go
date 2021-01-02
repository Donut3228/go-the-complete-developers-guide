package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
