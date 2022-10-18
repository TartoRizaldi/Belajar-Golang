package main

import "fmt"

func Penjumlahan(x int, y int) (pen int) {
	pen = x + y
	return
}
func Pembagian(x int, y int) (pem int) {
	pem = x / y
	return
}

func Home() (string, string) {
	return "Bandung", "Programer"
}

func main() {
	Alamat, Pekerjaan := Home()

	fmt.Println("ini Home:", Alamat, Pekerjaan)
	total := Penjumlahan(10, 3)
	fmt.Println("Penjumlaahan:", total)
	fmt.Println("==================")

	totall := Pembagian(2, 1)
	fmt.Println("Pemagian:", totall)
}
