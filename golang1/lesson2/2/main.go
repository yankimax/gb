package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Для вычисления длинны и диаметра окружности, введите площадь")
	var S float64
	fmt.Scan(&S)

	D := 2 * math.Sqrt(S/math.Pi)
	fmt.Println("Диаметр:", D)

	P := math.Pi * D
	fmt.Println("Длина окружности (периметр):", P)
}
