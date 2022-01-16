package main

import (
	"fmt"
	"os"
)

var fsave map[int]int

func main() {
	var num int
	fsave = make(map[int]int)

	for {
		fmt.Println("Введите число, для которого нужно расчитать число Фиббоначи")
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Допускается ввод только целого числа.")
			os.Exit(1)
		}
		fmt.Println("Ответ: ", fcalc(num))
	}
}

func fcalc(n int) int {
	_, ok := fsave[n]
	if !ok {
		if n < 2 {
			fsave[n] = n
		} else {
			fsave[n] = fcalc(n-1) + fcalc(n-2)
		}
	}
	return fsave[n]
}
