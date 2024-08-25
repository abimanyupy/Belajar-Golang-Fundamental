package main

import "fmt"

func main() {
	// defer func(){ //anonymous function
	// 	if r := recover(); r != nil{ //recover() ini akan mereturn error yg disebabkan panic || ketika tidak panic maka recover akan mereturn ni
	// 		fmt.Println("Terjadi Panic: ", r)
	// 	}
	// }()

	// //sebelum program berhernti karena kejadian panic akan menjalankan defer recover
	// //karena ada kejadian panic maka recover() akan mereturn nilai error yg ada didalam panic (func isEmpty) 
	// //sehingga nilai r adalah string yg ada didalam panic ("Inputan tidak boleh kosong!")

	var input string
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&input)

	fmt.Println("Start")
	if !isEmpty(input){
		fmt.Println(input)
	}
	fmt.Println("Done")
}

func isEmpty (input string) (empty bool){

	//jika program ingin berjalan ke line selanjutnya defer pindah ke sini
	//meskipun terjadi panic line selanjutnya tetap dijalankan
	defer func(){ 
		if r := recover(); r != nil{ 
			fmt.Println("Terjadi Panic: ", r)
			empty = true
		}
	}()

	if input == ""{
		panic("Inputan tidak boleh kosong!")
	}
	empty = false
	return 
}