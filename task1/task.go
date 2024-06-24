package main

import (
	"fmt"
	"os"
)

func main() {

	var a int
	var b int
	var sign string

	fmt.Println("Введите числа: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("Введите операцию: ")
	fmt.Fscan(os.Stdin, &sign)
	if sign == "+" {
		fmt.Println(a + b)
	} else if sign == "-" {
		fmt.Println(a - b)
	} else if sign == "*" {
		fmt.Println(a * b)
	} else if sign == "/" {
		fmt.Println(a / b)
	}
}
