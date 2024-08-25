package main

import "fmt"

func main() {
	//anonymous function
	func(){
		fmt.Println("anonymous function 1")
	}()

	//anonymous function dalam variabel
	halo := func(){
		fmt.Println("anonymous function 2")
	}
	halo()

	//passing argument ke dalam anonymous function
	func(halo string){
		fmt.Println(halo)
	}("anonymous function 3")

	//passing argument dalam variabel function
	hello := func(halo string){
		fmt.Println("Hello, ",halo)
	}
	hello("Abimanyu")

	fmt.Println("==========================================")

	//passing anonymous function sebagai argumen
	greetEnglish := func(name string) string {
		return "Hello " + name
	}
	greetRussian := func(name string) string {
		return "Привет " + name
	}
	greetIndo := func(name string) string {
		return "Hai " + name
	}
	greet("Abimanyu", greetEnglish)
	greet("Abimanyu", greetRussian)
	greet("Abimanyu", greetIndo)

	fmt.Println("==========================================")

	add := func(num1,num2 int) int {
		return num1 + num2
	}
	multiply := func(num1,num2 int) int {
		return num1 * num2
	}

	calculate(10,20,add)
	calculate(10,20,multiply)
}

func greet (name string, f func(name string) string) {
	fmt.Println(f(name))
}

func calculate(a,b int, operator func(x,y int) int) {
	fmt.Println(operator(a,b))
}