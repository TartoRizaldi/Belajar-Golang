package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Word")

}

func main() {

	for i := 0; i < 10; i++ {
		sayHello()
	}
	// // Array Biasa
	// var name [3]string
	// name[0] = "Rizal"
	// name[1] = "aldi"
	// name[2] = "tio"
	// fmt.Println(name[0], name[1], name[2])

	// var value = [3]int{
	// 	134,
	// 	121,
	// 	12,
	// }
	// fmt.Println(value)

	// //Slice
	// var Kos = [...]string{
	// 	"Dimarto",
	// 	"Rijal",
	// 	"SUA",
	// 	"bOD",
	// 	"eLBVI",
	// 	"EWSDS",
	// 	"WEOIDS",
	// 	"EWEWE",
	// 	"WEWE",
	// 	"DWDWEW",
	// }

	// var slice1 = Kos[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))

}
