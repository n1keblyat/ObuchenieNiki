package main

import "fmt"

func main() {
	numbers := make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)

	fmt.Println(numbers[0])
	fmt.Println(numbers[9])

	for a := 4; a <= 8; a++ {
		numbers[a] = a + 2
	}
	numbers[len(numbers)-1] = 0
	numbers = numbers[:len(numbers)-1]

	fmt.Println(numbers)

}
