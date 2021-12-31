package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Первое число должно быть целым или с плавающей точкой")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Второе число должно быть целым или с плавающей точкой")
		os.Exit(1)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя")
			os.Exit(1)
		}
		res = a / b
	case "pow":
		res = float32(math.Pow(float64(a), float64(b)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
