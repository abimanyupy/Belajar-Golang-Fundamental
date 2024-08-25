package main

import "fmt"

type Patient struct {
	Name string
	Age int
	Celcius
}

//bisa dibuat di package / folder sendiri utk menampung custom type
type Celcius float64  //custom type Celcius dengan tipe float64 dan huruf besar agar exported
//tujuannya agar kita dapat membuat tipe data yang lebih spesifik agar lebih mudah dipahami

func (c Celcius) tofarenheit() float64 {  //value reciver
	return float64((c * 9 / 5) + 32) //agar return menjadi float64
}

func (c *Celcius) add(value float64) {
	*c += Celcius(value)  //merubah value yg tipenya float64 menjadi Celcius
	
}

func main(){
	var temperature Celcius = 20.0 //variabel temperature bertipe data Celcius
	fmt.Println("Temperature :",temperature)
	fmt.Println("Suhu di ruangan ini :", temperature.tofarenheit(), "derajat Farenheit")
	temperature.add(3)
	fmt.Println("Suhu di ruangan ini menjadi :", temperature, "derajat Celcius")

	fmt.Println("===============================================")

	newPatient := Patient{
		Name: "Abimanyu",
		Age: 23,
		Celcius: 39.0,
	}
	fmt.Printf("newPatient : %+v \n",newPatient)
	fmt.Println(newPatient)
}