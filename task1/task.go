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
	_, err = fmt.Fscan(os.Stdin, &sign)
	if err != nil {
		log.Fatal(errors.New("введен не знак"))
	}
	calculator(a, b, sign)
}

func calculator(a, b int, sign string) {
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
