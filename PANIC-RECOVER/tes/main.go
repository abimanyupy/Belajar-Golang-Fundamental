package main

import "fmt"

func main() {
	defer func(){
		fmt.Println("Deffered Function")
	}()

	fmt.Println("Before panic")
	panic("panic occured")

}