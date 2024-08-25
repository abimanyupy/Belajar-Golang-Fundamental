package main

import "fmt"

func main() {
	var buah = [4]string{"Apel", "Jeruk", "Anggur", "Semangka"}

	fmt.Println(buah, len(buah), buah[3])

	buah[3] = "Melon"
	fmt.Println(buah)

	if buah[3] == "Melon" {
		buah[3] = "Mengkudu"
		fmt.Println(buah)
	}

	fmt.Println("=======================================")

	var nilai [3]int //array dengan panjang 3 jika ingin tanpa panjang array kasih [...]
	nilai[0] = 10
	nilai[1] = 20
	nilai[2] = 30
	fmt.Println(nilai)

	fmt.Println("=======================================")

	var makanan = [...]string{"roti", "nasi", "buah", "permen"}
	fmt.Println(makanan)
	fmt.Println("jumlah makanan tersedia :", len(makanan))

	fmt.Println("=======================================")
	//kombinasi array dengan looping
	fmt.Println("daftar makanan")
	for i := 0; i < len(makanan); i++ {
		fmt.Printf("%d. %s \n", i+1, makanan[i]) //i+1 agar angka yg muncul dari 1 bukan 0
	}

	fmt.Println("=======================================")

	for i, makan := range makanan { //makan adalah variabel yg merepresentasikan tiap element yg ada dialam array makanan
		fmt.Printf("%d. %s \n", i+1, makan)
	}

	fmt.Println("=======================================")
	//looping tanpa index
	for _, makan := range makanan {
		fmt.Printf("Nama menu : %s \n", makan)
	}

	fmt.Println("=======================================")

	for i := range makanan {
		fmt.Printf("Index menu ke : %d \n", i)
	}

	fmt.Println("=======================================")
	//looping array dengan pengkondisian
	var angka = [5]int{3, 8, 2, 7, 4}

	for i := 0; i < len(angka); i++ {
		if angka[i]%2 == 0 {
			fmt.Println("angka", angka[i], "genap")
		} else {
			fmt.Println("angka", angka[i], "ganjil")
		}
	}

	fmt.Println("=======================================")
	//cetak angka genap
	var numbers = [6]int{3, 8, 2, 7, 4, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}

	fmt.Println("=======================================")
	//perkalian array
	fmt.Println("Sebelumnya :", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
	fmt.Println("Sesudahnya :", numbers)

	fmt.Println("=======================================")
	//array multidimensi
	var matrix = [2][3]int{
		{1, 2, 3},  //ini array index 0
		{4, 5, 6},  //ini array indes 1
	} //membuat array multidimensi 2 baris 3 kolom

	fmt.Println(matrix)
	fmt.Println(matrix[1][1]) //mengakses array multidimensi index 1 kolom ke 1 yang berisi angka 5
	fmt.Println(matrix[1][2])

}
