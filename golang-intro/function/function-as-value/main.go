package main

import "fmt"

//func as value adalah kemampuan utk menentapkan sebuah variabel yg memiliki value sebuah function
func main() {
	//membuat variabel sum dgn tipe sebuah function yg menerima 2 parameter integer dan returnnya integer dan valuenya diisi dengan func add
	//sehingga dapat memberlakukan sum seperti function biasa, krn sum ini tipenya function
	var sum func(int, int) int = add //harus disesuaikan dgn func add yg sdh dibuat utk tipenya
	fmt.Println("hasil :",sum(10, 20))
}

func add(a, b int) int {
	return a + b
}