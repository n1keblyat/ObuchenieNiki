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
	_, err = fmt.Fscan(os.Stdin, &b)
	if err != nil {
		log.Fatal(errors.New("введено не число"))
	}
	fmt.Println("Введите операцию: ")

	fmt.Fscan(os.Stdin, &sign)

	_, err = fmt.Fscan(os.Stdin, &sign)
	if err != nil {
		log.Fatal(errors.New("введен не знак"))
	}
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
