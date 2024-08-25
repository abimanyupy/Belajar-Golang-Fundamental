package main

import "fmt"

func main() {
	numberSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numberSlice)
	fmt.Println("panjang slice:",len(numberSlice))
	fmt.Println("kapasitas slice:",cap(numberSlice))
	
	fmt.Println("===================================================")

	scores := make([]int, 4, 7) //make() menerima 3 argumen yaitu tipeSlice, panjang, kapasitas(opsional ketika menggunakan make())
	scores[0], scores[1], scores[2], scores[3] = 10, 20, 30, 40
	fmt.Println(scores)
	fmt.Println("panjang scores:",len(scores))
	fmt.Println("kapasitas scores:",cap(scores))

	fmt.Println("===================================================")

	heroes := [4]string{"Spiderman", "Superman", "Batman", "Ironman"}
	fmt.Println("heroes:",heroes)
	// heroes[5] = "Hulk" //bakal error tidak bisa nambahin isi array

	villain := make([]string, 3, 5)
	villain[0],villain[1],villain[2] = "Joker","Harley","Deadpool"
	fmt.Println("villain:",villain)
	fmt.Println("panjang villain:",len(villain))
	fmt.Println("kapasitas villain:",cap(villain))

	villain = append(villain, "Hulk") //append(namaSlice, value) bisa untuk menambahkan lebih dari 1 value
	villain = append(villain, "ultron", "Captain")
	fmt.Println("villain:",villain)
	fmt.Println("panjang villain:",len(villain))	
	fmt.Println("kapasitas villain:",cap(villain)) //kapasitas slice akan di x2 ketika elemen sudah melebihi batas kapasitas 

	fmt.Println("===================================================")

	var numbers = [4]int{2, 1, 6, 8}
	var anotherNumbers = numbers //ini merupakan pass by value //array numbers diduplikasi menjadi anotherNumbers sehingga array yg berbeda
	fmt.Println("numbers:",numbers)
	fmt.Println("anotherNumbers:",anotherNumbers)
	anotherNumbers[0] = 10
	fmt.Println("numbers:",numbers)
	fmt.Println("anotherNumbers:",anotherNumbers)

	fmt.Println("===================================================")

	var harga = []int{1000, 2000, 3000, 4000}
	var anotherHarga = harga //ini merupakan pass by reference //slice yg ada didalam variabel harga akan disimpan referensinya didalam variabael anotherHarga sehingga ketika merubah elemen di var anotherHarga maka akan menyimpan referensi yg sama dengan var harga
	fmt.Println("harga:",harga)
	fmt.Println("anotherHarga:",anotherHarga)
	anotherHarga[0] = 5000 //kode ini akan mengganti array anotherHarga [0] menjadi 5000 dan array harga juga akan ikut berubah pada array harga[0] menjadi 5000
	fmt.Println("harga:",harga)
	fmt.Println("anotherHarga:",anotherHarga)

	fmt.Println("===================================================")
	//membuat slice dari array yg sudah ada

	umur := [4]int{20, 21, 22, 23}
	sliceUmur := umur[0:2] //membuat slice baru dgn nama sliceUmur dari array umur yg isi slicenya adalah dari index 0 array umur sejumlah 3 elemen
	//umur[0:3] 0 adalah start yg mau diambil mulai dari index ke brp dan 3 adalah jumlah elemen yg mau diambil
	//akan ada kejadian pass by reference jadi ketika value slice diubah maka value array juga akan berubah
	fmt.Println("umur:",umur)
	fmt.Println("sliceUmur:",sliceUmur)

	sliceUmur[0] = 10
	fmt.Println("umur:",umur)
	fmt.Println("sliceUmur:",sliceUmur)

	sliceUmur = append(sliceUmur, 24)
	fmt.Println("umur:",umur)
	fmt.Println("sliceUmur:",sliceUmur)

}