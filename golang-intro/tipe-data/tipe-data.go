package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")

	var positiveNumber uint8 = 90 //uint hanya bisa positif
	var negativeNumber int = -90 //int bisa untuk positif dan negatif

	fmt.Printf("Bilangan Positif : %d \n", positiveNumber) //%d untuk format tipe data bilangan bulat
	fmt.Printf("Bilangan negatif : %d \n", negativeNumber)

	var decimalNumber = 3.90

	fmt.Printf("Bilangan Desimal : %f \n", decimalNumber)  //%f untuk format tipe data bilangan pecahan
	fmt.Printf("Bilangan Desimal : %.2f \n", decimalNumber) //.2 untuk cetak 2 angka dibelakang koma

	fmt.Println("=====================================================")

	fmt.Println("Boolean")
	var exist = true
	fmt.Printf("exist ? %t \n", exist) //%t unut format tipe data boolean

	fmt.Println("=====================================================")

	fmt.Println("String")
	var message string = "Hallo"
	fmt.Printf("message : %s \n", message) //%s unut format tipe data string

	var otherMessage = `Nama saya "Abimanyu", 
Salam Kenal, 
Mari belajar di "Enigma Camp" `
	fmt.Println(otherMessage)

	fmt.Println("Abimanyu" + "Putro")

	fmt.Println(len(message))

}