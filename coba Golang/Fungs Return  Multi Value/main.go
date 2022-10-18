package main

import "fmt"

func Data() (string, string, int) {

	return "Tarto Rizaldi", "Back End Developer", 22

}

func main() {
	NamaLengkap, Pekerjaan, Umur := Data()

	fmt.Println(NamaLengkap, Pekerjaan, Umur)
}
