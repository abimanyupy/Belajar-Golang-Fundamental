package main

import "fmt"

var nama = "abimanyu" //scope global

func main() { //func namaFunction(parameter1 tipeData, parameter2 tipeData) tipeData{body function}
	//func main berguna sbg entry point/yg pertama jalan

	hello()
	fmt.Println("dari main :", nama)
	salam("abimanyuu", 23)
	add(7, 2)

	total := add(8, 4)
	fmt.Println("ini total :", total)

	if total > 10 {
		fmt.Println("ini lebih dari 10")
	}

}

func hello() {
	var nama = "putro" //scope local/variabel lokal //pada golang apabila ada nama var yg sama maka yg digunakan yg lokal dlu
	fmt.Println("Hello World!")
	fmt.Println("dari hello :", nama)

}

func salam(nama string, umur int) {
	fmt.Println("Salam Kenal saya :", nama, "umur :", umur)
}

func add(a int, b int) int {
	fmt.Println("a + b = ", a+b)
	return a + b
}
