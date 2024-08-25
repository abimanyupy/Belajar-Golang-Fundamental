package main

import (
	"fmt"
)

func main() {

	var firstName string = "abimanyu"

	var lastName = "putro"

	fmt.Println(firstName, lastName)
	fmt.Printf("test\n") // \n berfungsi sebagai enter pada Printf untuk membuat baris baru
	fmt.Printf("Halo, %s spasi %s  \n", firstName, lastName)

	var (
		age     int
		address string
	)

	age = 23
	address = "Ungaran"

	fmt.Printf("Halo, saya %s %s berumur %d dan saya tinggal di %s \n", firstName, lastName, age, address)

	var (
		bootcampName, bootcampAddress, bootcampNumber = "Enigma Camp", "Jakarta Selatan", 1
	)
	fmt.Println(bootcampName, bootcampAddress, bootcampNumber)

	day := "Selasa"
	date := "27"
	month := "maret"

	fmt.Println(day, date, month+" 2024")

	var number = 10
	number = 222

	const phi = 3.14

	fmt.Println(number, phi)

	animal, _ := "kucing", "anjing"
	// fmt.Println(animal)
	fmt.Scanln(&animal)
}
