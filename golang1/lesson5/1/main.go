package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Println("Введите число, для которого нужно расчитать число Фиббоначи")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Допускается ввод только целого числа.")
		os.Exit(1)
	}
	fmt.Println("Ответ: ", fcalc(num))
}

func fcalc(n int) int {
	if n < 2 {
		return n
	}
	return fcalc(n-1) + fcalc(n-2)
}
