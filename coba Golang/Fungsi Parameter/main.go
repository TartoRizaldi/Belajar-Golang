package main

import "fmt"

func Nama(FirstName string, LastName string) {
	fmt.Println("Nama:", FirstName, LastName)

}

func main() {
	Alamat := "Bekasi"
	Nama(Alamat, "Rijaj	")
	Nama("Tarto", "Rizaldi")

}
