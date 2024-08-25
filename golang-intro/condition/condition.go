package main

import "fmt"

func main() {
	if false {
		fmt.Println("Kode dijalankan")
	}

	fmt.Println("Done")

	fmt.Println("==================================================")

	result := "gagal" //jika result = lulus maka akan jalankan line 17 jika result = gagal maka akan jalankan line 19

	if result == "lulus" {
		fmt.Println("Selamat Anda,", result)
	} else {
		fmt.Println("Maaf Anda,", result)
	}

	fmt.Println("==================================================")

	//deklarasi variabel temporary hanya dapat menggunakan type inverens(:=) jika menggunakan var maka akan error

	if result := "lulus"; result == "lulus" {  //jika variabel dideklarasikan didalam if statment maka hanya dapat diakses oleh statment tsb.
		fmt.Println("Selamat Anda,", result)
	} else {
		fmt.Println("Maaf Anda,", result)
	}

	fmt.Println("==================================================")

	hour := 10

	if hour > 8 && hour < 15 {
		fmt.Println("Masih jam kerja")
	} else {
		fmt.Println("Diluar jam kerja")
	}

	fmt.Println("==================================================")

	IPK := 1.3
	var graduationStatus string

	if IPK >= 3.5 && IPK <= 4.0 {
		graduationStatus = "Cumlaude"
	} else if IPK >= 3.0 && IPK < 3.5 {
		graduationStatus = "Memuaskan"
	} else if IPK >= 2.5 && IPK < 3.0 {
		graduationStatus = "Cukup Memuaskan"
	} else {
		graduationStatus = "Kurang memuaskan"
	}

	fmt.Printf("Selamat Anda lulus dengan predikat %s dengan IPK %.2f", graduationStatus, IPK)

	fmt.Println("==================================================")

	x := -1
	y := 2

	if x > 0{
		if y > 0 {
			fmt.Println("x > 0 dan y > 0")
		} else {
			fmt.Println("x > 0 dan y <= 0")
		}
	} else {
		if y > 0 {
			fmt.Println("x < 0 dan y > 0")
		} else {
			fmt.Println("x < 0 dan y < 0")
		}
	}

	fmt.Println("==================================================")

	var poin = 10
	// switch poin {
	// 	case 8: {
	// 		fmt.Println("Bagus")
	// 	}
	// 	case 7, 6, 5: {
	// 		fmt.Println("Cukup")
	// 	}
	// 	default: {
	// 		fmt.Println("Kurang")
	// 	}
	// }
	switch {
		case poin == 10: {
			fmt.Println("Sempurna")
		}
		fallthrough //ini digunakan untuk mengeksekusi statement selanjutnya tanpa menghiraukan case yang ada meskipun FALSE tetap dieksekusi
					//hanya berlaku untuk 1 case saja
		case poin >= 8 && poin < 10: {
			fmt.Println("Bagus")
		}
		case poin >= 7 && poin < 8: {
			fmt.Println("Cukup")
		}
		default: {
			fmt.Println("Nilaimu kurang memuaskan")
			fmt.Println("Anda perlu mengulang ujian")

		}
	}
}
