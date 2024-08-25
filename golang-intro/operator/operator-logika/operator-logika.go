package main

import "fmt"

func main() {
	fmt.Println("Operator Logika")

	a:= true
	b:= false
	fmt.Println(a && b) //AND jika salah satu ada yg false return akan false(harus true semua agar return true)
	fmt.Println(a || b) //OR akan return true ketika salah satu nilai ada yg true dan jika keduanya false maka return false
	fmt.Println(!a)
	fmt.Println(false || true && false)
}