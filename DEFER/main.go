package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("Done") //defer akan dieksekusi di akhir func meskipun di line awal/tengah/akhir  //dipanggil tapi ditunda utk diakhir func
	// fmt.Println("Start")
	// fmt.Println("Processing")

	// num := 4

	// defer fmt.Println("Result defer :", num) //tetap mengambil num tepat diatasnya namun akan dieksekusi di akhir func

	// num += 10
	// num *= 2

	// fmt.Println("Result main:", num)

	// fmt.Println("==================================================")

	// num1 := 4
	// num2 := 0

	// defer fmt.Println("Done")
	// fmt.Println("Start")
	// fmt.Println(num1 / num2)

	// fmt.Println("==================================================")

	// defer add(7, multiply(2,3))  //yg di defer cuma statment add saja || sedangkan statment multiplay tetap dieskekusi terlebih dahulu
	// add(3, 4)

	// fmt.Println("==================================================")

	// fmt.Println("Start")
	// defer fmt.Println("Done")
	// defer add(1,2)
	// defer add(3,4)
	// defer multiply(3,3)
	// //outputnya akan start 9 7 3 Done
	// //jika ada lebih dari 1 defer maka akan LIFO(last in first out) cetak defer yg akhir dulu atau dari bawah ke atas

	// fmt.Println("==================================================")

	fmt.Println("Start")
	defer loop()
	fmt.Println("Done")

}

func add(num1, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}

func multiply(num1, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}

func loop() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done Loop")
}
