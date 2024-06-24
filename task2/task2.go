package main

import "fmt"

func main() {
	numbers := make([]int, 0, 10)
	for i := 1; i < 11; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

}
