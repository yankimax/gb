package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var a int64
	fmt.Println("Введите целое число:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Число должно быть целым!")
		os.Exit(1)
	}
	if big.NewInt(a).ProbablyPrime(0) {
		fmt.Println("Число ", a, " простое")
		os.Exit(1)
	}
	fmt.Println("Число ", a, " не простое")
}
