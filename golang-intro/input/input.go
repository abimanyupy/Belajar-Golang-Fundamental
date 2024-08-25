package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var firstName string
	var lastName string
	var age int

	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("Masukkan Umur : ")
	fmt.Scanln(&age)
	birthYear := (2024 - age)

	fmt.Println("Your Name is", firstName, lastName, "You are born in", birthYear)

	fmt.Println("==================================================================")

	var fullName string
	var address string
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data Diri Saya Sendiri")
	fmt.Print("Masukkan Nama Lengkap : ")
	scanner.Scan() //untuk mendapatkan input dari user
	fullName = scanner.Text() //value dapat disimpan dalam fullName dan utk mendpatkannya menggunakan scanner.Text()

	fmt.Print("Masukkan Umur Anda : ")
	scanner.Scan() //utk mendapatkan input dari user
	age, _ = strconv.Atoi(scanner.Text()) //strconv maksudnya adalah string conversion dan Atoi adalah Ascii to integer

	fmt.Print("Masukkan Alamat Anda : ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("%s | %d | %s \n", fullName, age, address)
}
