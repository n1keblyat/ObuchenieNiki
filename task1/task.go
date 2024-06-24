package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	var a int
	var b int
	var sign string

	fmt.Println("Введите числа: ")
	_, err := fmt.Fscan(os.Stdin, &a)
	if err != nil {
		log.Fatal(errors.New("введено не число"))
	}
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("Введите операцию: ")
	fmt.Fscan(os.Stdin, &sign)
	switch sign {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		log.Fatal(errors.New("переданной операции не существует"))
	}
}
