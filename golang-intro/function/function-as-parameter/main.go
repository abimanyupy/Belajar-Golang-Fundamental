package main

import "fmt"

//function juga dapat dijadikan paramter
func main() {
	numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //membuat slice baru dgn nama numbers

	evenNumbers := filter(numbers, isEven)  //disini karena memberikan argumennya isEven artinya f pada func filter bagian (( if f(num){result = append(result, num)} )) dapat diartikan sebagai func isEven itu sendiri
	//membuat variabel baru dgn nama evenNumbers utk menampung hasil return dari filter
	//memanggil func filter dgn argumen1 numbers dan argumen2 isEven
	fmt.Println("Genap :", evenNumbers)

	oddNumbers := filter(numbers, isOdd)
	fmt.Println("Ganjil :", oddNumbers)
}

func filter(numbers []int, f func(int) bool) []int {
	//func f ini akan mereturn boolean dan ketika func f dipanggil nilai boolean tsb yg akan menentukan apakah num akan diappend ke result atau tidak
	var result []int //membuat variabel baru dgn nama result 

	for _, num := range numbers { //for ini berguna utk pengecekan dgn pemanggilan func f yg dari parameter
		if f(num) { //bagian ini dapat diartikan sbg isEven/isOdd karena memberikan argumen isEven/isOdd pada func filter pada func main
					//num ini adalah setiap elemen yg ada di slice numbers (func main)
					//jadi setiap elemen numbers yg bernilai true saat pemanggilan func f ini akan dilakukan append num tsb ke slice result
			result = append(result, num)
		}
	}

	return result //hasil result ini akan ditamung pada evenNumbers/oddNumbers
}

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}
