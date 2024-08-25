package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Angkat", i)
	}

	fmt.Println("============================================")

	var i = 6
	for i <= 10 {
		fmt.Println("Angka", i)
		i++ //perlu diberikan i++ agar berhenti jika tidak maka  akan infiniti loop karena true terus tidak melebihi 10
	}

	fmt.Println("============================================")

	var index = 1
	for {
		fmt.Println("Angka", index)
		index++
		if index == 5 {
			break
		}
	}

	fmt.Println("============================================")
	//nested looping
	for i := 1; i <= 3; i++ { //menentukan banyak barisnya
		for j := 1; j <= 6; j++ { //menentukan jumlah data yg dicetak berulang tiap barisnya
			fmt.Print(j, " ")
		}
		fmt.Println() //berguna untuk berpindah baris baru
	}

	fmt.Println("============================================")
	//break berguna utk keluar dari looping

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 { //jika perulangan sudah sampai 5 akan keluar dari for dan lanjut ke line 59
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("Finish")

	fmt.Println("============================================")
	//continue berguna utk lanjut ke perulangan selanjutnya tanpa eksekusi kode yg tersisa di looping

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Print("(Skip 5) ")
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("Finish")

	fmt.Println("============================================")

	fmt.Println("Start")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 4 {
				// break
				continue
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	fmt.Println("Finish")

}
