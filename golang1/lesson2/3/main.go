package main

import "fmt"

func main() {
	fmt.Println("Введите трехзначное число:")
	var Num int
	fmt.Scan(&Num)

	var a = Num % 10
	fmt.Println("В числе ", Num, " - ", a, " единиц")

	b := Num / 10 // от 321 оставляем 32
	b = b % 10
	fmt.Println("В числе ", Num, " - ", b, " десятков")

	c := Num / 100 // от 321 оставляем 3
	fmt.Println("В числе ", Num, " - ", c, " сотен")
}
