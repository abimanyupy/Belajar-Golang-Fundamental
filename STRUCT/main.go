package main

import "fmt"

type kendaraan struct {
	merk  string
	tahun int
	model string
	harga int
}

func main() {
	var a kendaraan
	a.merk = "Toyota"
	a.tahun = 2022
	a.model = "Avanza"
	a.harga = 2000000

	fmt.Println(a)
	fmt.Println(a.merk)
	fmt.Println(a.tahun)
	fmt.Println(a.model)
	fmt.Println(a.harga)
	fmt.Printf("tes %s", a.merk)

	fmt.Println("==================================")

	var b = kendaraan{
		merk:  "Toyota B",
		tahun: 2022,
		model: "Avanza B",
		harga: 2000000,
	}

	fmt.Println(b)
	fmt.Println(b.merk)
	fmt.Println(b.tahun)
	fmt.Println(b.model)
	fmt.Println(b.harga)

	fmt.Println("==================================")

	var c = kendaraan{}

	c.merk = "Toyota C"
	c.tahun = 2022
	c.model = "Avanza C"
	c.harga = 2000000

	fmt.Println(c)
	fmt.Println(c.merk)
	fmt.Println(c.tahun)
	fmt.Println(c.model)
	fmt.Println(c.harga)

	fmt.Println("==================================")

	var x = kendaraan{
		merk:  "Honda",
		tahun: 2023,
		model: "HRV",
		harga: 22000000,
	}

	var y kendaraan = x

	fmt.Println("x :", x)
	fmt.Println("y :", y)
	fmt.Printf("alamat memori x : %p\n ", &x)
	// fmt.Printf("alamat memori y : %p\n ", y)
	y.model = "Jazz"
	fmt.Println("x :", x)
	fmt.Println("y :", y)

	// fmt.Println("==================================")
	
	// var z = kendaraan {
	// 	merk:  "Honda",
	// 	tahun: 2029,
	// 	model: "Civic",
	// 	harga: 55000000,
	// }

	// updateKendaraan(z)
	// fmt.Println("kendaraan di function main :", z)

	// fmt.Println("==================================")

	// var z2 *kendaraan = &z
	// fmt.Printf("alamat memori z : %p\n ", &z)
	// fmt.Printf("alamat memori z2 : %p\n ", z2)
	// z2.model = "CRV"
	// fmt.Println("z :", z)
	// fmt.Println("z2 :", z2)

	fmt.Println("==================================")
	var z = kendaraan{
		merk:  "Hondaa",
		tahun: 2032,
		model: "Civicc",
		harga: 555000000,
	}
	updateKendaraan(&z)
	fmt.Println("Kendaraan di function main :", z)

	fmt.Println("==================================")

	var k = new(kendaraan)
	fmt.Printf("alamat memori k : %p\n ", k) //tidak perlu & untuk menjadikan pointer karena new(kendaraan) membuat k menjadi pointer kendaraan

}

func updateKendaraan(newKendaraan *kendaraan){
	newKendaraan.merk = "Toyota"
	newKendaraan.model = "Camry"
	newKendaraan.harga = 750000000
	newKendaraan.tahun = 2020
	fmt.Println("Kendaraan dari function updateKendaraan :", newKendaraan) 
}
