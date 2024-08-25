package main

import (
	"fmt"
)

func main() {
	// var angka1, angka2 int
	// // var angka2 int
	// fmt.Println("masukkan angka pertama")
	// fmt.Scanln(&angka1)
	// fmt.Println("masukkan angka kedua")
	// fmt.Scanln(&angka2)
	// fmt.Println("angka pertama awal", angka1)
	// fmt.Println("angka kedua awal", angka2)
	// fmt.Println("====================================")

	// angka1 = angka1 + angka2
	// angka2 = angka1 - angka2
	// angka1 = angka1 - angka2
	// fmt.Println("angka pertama akhir", angka1)
	// fmt.Println("angka kedua akhir", angka2)


	var nama string
	var age int
	fmt.Println("masukkan nama")
	fmt.Scanln(&nama)
	fmt.Println("masukkan umur")
	fmt.Scanln(&age)
	fmt.Printf("Hello %s, you are %d years old", nama, age)




}


