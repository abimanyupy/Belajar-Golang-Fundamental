package main

import "fmt"

func main () {
	fmt.Println("Operator Aritmatika")
	x := 9
	y := 3
	fmt.Println(x+y)
	fmt.Println(x-y)
	fmt.Println(x*y)
	fmt.Println(x/y)
	fmt.Println(x%y)
	fmt.Println(2 + 8 * 3 - 5)

	if condition := x % y == 0; condition {
		fmt.Println("YA")
	} else {
		fmt.Println("TIDAK")
	}

}