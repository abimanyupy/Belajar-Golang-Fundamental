package main

import (
	"errors"
	"fmt"
)

func main(){
	result, err := divide(1,0)

	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("Hasil pembagian: ",result)
}

func divide (a, b int) (int, error){
	if a == 0 || b == 0 {
		return 0, errors.New("pembagian dengan angka 0")
	}
	return a / b, nil
}