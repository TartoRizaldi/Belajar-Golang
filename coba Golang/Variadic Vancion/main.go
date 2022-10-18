package main

import "fmt"

func SumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func MinAll(No ...int) int {
	totall := 0
	for _, value := range No {
		totall -= value
	}
	return totall
}

func main() {
	total := SumAll(10, 20, 30, 40, 50)
	fmt.Println("Percobaan Pertama:", total)

	slice := []int(20, 2, 20, 2, 50)
	total = SumAll(slice...)
	fmt.Println(total)

	totall := MinAll()
	fmt.Println("Percobaan Kedua:", totall)
}
