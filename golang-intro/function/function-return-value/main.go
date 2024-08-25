package main

import "fmt"

func main() {
	hasil := add(10, 1)
	fmt.Println("hasil :", add(10, 20))
	fmt.Println("hasil cara 2:", hasil)

	hasil2 := multiply(10, 2)
	fmt.Println("hasil perkalian :", hasil2)

	total := add(10, multiply(2, 3))
	fmt.Println("total1 :", total)

	total2 := add(10, 1)
	total3 := multiply(total2, 2)
	fmt.Println("total3 :", total3)

}

func add(a int, b int) int { //int itu tipe data dari return value
	result := a + b
	return result
}

func multiply(a, b int) int { //cara penulisan tipe data jika sama
	result := a * b
	return result
}