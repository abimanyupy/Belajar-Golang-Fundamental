package main

import "fmt"

//variadic function adalah func di golang yg dapat menerima banyak argumen tanpa batas
//variadic ini bertipe slice ([]tipeData)
func main() {
	total := sum(1, 9, 8, 7)
	fmt.Println("total :", total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
