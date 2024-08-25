package main

import "fmt"

type Shape interface {
	luas() float64
	keliling() float64
}

//bisa dibilang struc rectangle dan square mengimplementasikan interface shape krn memiliki method yang sama dari interface shape

type Rectangle struct {
	width, height float64
}

func (r Rectangle) luas() float64 {
	return r.width * r.height
}

func (r Rectangle) keliling() float64 {
	return 2 * (r.width + r.height)
}

type Square struct { //sturct yang mengimplementasikan semua method yang ada di dalam interface
	side float64
}

func (s Square) luas() float64 {
	return s.side * s.side
}

func (s Square) keliling() float64 {
	return 4 * s.side
}

// func luasRectangle(r Rectangle) {
// 	fmt.Println("luas : ", r.luas())
// }

// func luasSquare(s Square) {
// 	fmt.Println("luas : ", s.luas())
// }
//2 function diatas dapat dipersingkat dgn merubah parameternya menjadi interface shape

func luas(s Shape) {
	fmt.Println("luas : ", s.luas())
}

func keliling(s Shape) {
	fmt.Println("keliling : ", s.keliling())
}

func main() {
	var shape1 Shape = Square{side: 5} //contoh upcasting //upcasting ini hanya dapat dilakukan oleh struct yg sudah mengimplementasikan suatu interface
	var shape2 Shape = Rectangle{width: 4, height: 3}
	var shape3 Shape = Rectangle{width: 7, height: 9}

	fmt.Printf("shape1 : %#v\n", shape1) //%#v digunakan untuk mendapatkan nama type struct dan nama fieldnya
	fmt.Printf("shape2 : %#v\n", shape2)
	fmt.Printf("shape3 : %#v\n", shape3)

	fmt.Println("==================================================================")

	shapes := []Shape{shape1, shape2, shape3}
	for _, shape := range shapes {
		// fmt.Println("luas : ", shape.luas())
		// fmt.Println("keliling : ", shape.keliling())
		fmt.Printf("%#v memiliki luas %.1f dan keliling %1.f \n", shape, shape.luas(), shape.keliling())
	}

	fmt.Println("==================================================================")

	// luasRectangle(Rectangle{width: 3, height: 7})
	// luasSquare(Square{side: 9})

	luas(Rectangle{width: 2, height: 3})
	luas(Square{side: 7})

	keliling(Rectangle{width: 2, height: 3})
	keliling(Square{side: 7})

	fmt.Println("==================================================================")

	var square1 Shape = Square{side: 6}
	fmt.Println("Luas : ", square1.luas())
	fmt.Println("Keliling : ", square1.keliling())
	// fmt.Println("Side : ", square1.(Square).side) //bisa utk akses side
	// fmt.Println("Side : ", square1.side) //error

	var originalSquare Square = square1.(Square) //ini namanya downcasting //merubah variabel suqare1 yg awalnya tipenya Shape menjadi tipe Square //merubah menjadi original structnya
	fmt.Println("Luas : ", originalSquare.luas())
	fmt.Println("Keliling : ", originalSquare.keliling())
	fmt.Println("Side : ", originalSquare.side)

	fmt.Println("==================================================================")
	//conoth dynamic typing //interface kosong
	//bisa memberikan value apapun pada variabel yg tipe datanya interface kosong
	//interface kosong bisa upcasting dan downcasting

	// var anything interface{}
	var anything any

	//interface{} bisa diubah menjadi any --- var anything any

	anything = "hello"
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = 10
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = true
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = []string{"hello", "world"}
	fmt.Printf("anything %T: %v \n", anything, anything)

}
