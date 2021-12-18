package main

import "fmt"

func main() {
	var a, b, S float32

	fmt.Println("Для расчета площади треугольника введите два числа через пробел")
	fmt.Scan(&a, &b)

	S = a * b
	fmt.Println(S)
}
