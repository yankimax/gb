package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Пожалуйста, посмотрите исправление прошлой домашней работы (добавил свою реализацию нахождения простых чисел)
// https://github.com/YaninMaxim/gb/pull/3

func main() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		arr = append(arr, int(num))
	}

	sort(arr)

	for _, val := range arr {
		fmt.Println(val)
	}
}

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
