package main

import "fmt"

type kendaraan struct {
	merk  string
	tahun int
	model string
	harga int
	mesin
}

type mesin struct {
	tipe string
	cc int
}

func main() {
	var a = kendaraan{
		merk:  "Toyota",
		tahun: 2022,
		model: "Avanza",
		harga: 2000000,
		mesin: mesin{
			tipe: "matic",
			cc:  1600,
		},
	}
	fmt.Println("a: ", a)
	fmt.Println("a mesin.tipe: ", a.mesin.tipe)


}