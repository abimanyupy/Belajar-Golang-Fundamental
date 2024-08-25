package main

import "fmt"

type triangle struct {
	base float64
	height float64
}

func (t triangle) area() { //method area() ini dimiliki oleh struct triangle //ini adalah contoh value receiver(jadi mengambil salinna triangle utk jadi parameternya)
	fmt.Println("Luas segitiga : ", 0.5 * t.base * t.height)
}

func (t *triangle) increaseSize() { //harus memiliki pointer receiver
	t.base += 2
	t.height += 2
}

func main(){
	instanceTriangle := triangle{10,12}
	instanceTriangle.area()

	fmt.Println("==========================================")

	fmt.Println("instanceTriangle : ", instanceTriangle)
	instanceTriangle.increaseSize()
	fmt.Println("instanceTriangle : ", instanceTriangle)
	instanceTriangle.area()


}

//kpn pake value receiver dan kapan pake pointer receiver
//jika kita hanya ingin membaca value fieldnya saja kita pake value receiver
//jika inigin merubah nilai fieldnya kita pake pointer receiver


// //contoh method dgn return value
// func (t triangle) area() float64 {
// 	return 0.5 * t.base * t.height
// }

// func main(){
// 	instanceTriangle := triangle{10,12}
// 	area := instanceTriangle.area()
// 	fmt.Println("Luas segitiga : ", area)
// }