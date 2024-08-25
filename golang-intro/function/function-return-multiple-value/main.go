package main

import "fmt"

func main() {
	num := []int {123, 5, 6, 7, 8, 9, 10}

	kecil, besar := minMax(num)
	fmt.Println("kecil :", kecil)
	fmt.Println("besar :", besar)

	fmt.Println("===================================================")
	_, besarr := minMaxx(num) //jika ingin mencetak yg besar saja
	// fmt.Println("kecil :", kecill)
	fmt.Println("besar :", besarr)
}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return min, max
}


//contoh penggunaan named return value
func minMaxx(numbers []int) (min int, max int) { 
	min = numbers[0] //tinggal reassign karena sudah dideklarasikan diatas
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return //tidak perlu tulis min dan max karena sudah ditetapkan diatas ingin mengambalikan variabel min dan max
}