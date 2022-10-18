package main

import "fmt"

func NamedValue() (Nama string, Umur int, Pendidikan string, Pekerjaan string) {
	Nama = "Tarto Rizaldi, S.Kom"
	Umur = 22
	Pendidikan = "Sarjana"
	Pekerjaan = "Back End Developer"
	return
}

func main() {
	Nama, Umur, Pendidikan, Pekerjaan := NamedValue()
	fmt.Println("Nama:", Nama, "Umur:", Umur, "Pendidikan:", Pendidikan, "Pekerjaan", Pekerjaan)
}
