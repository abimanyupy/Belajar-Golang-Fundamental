package main

import (
	"access-modifier/library"  //import 1 pacakage dari folder libarary
	"fmt"
)

func main() {
	fmt.Println("HourInADay = ", library.HourInADay) //hanya bisa HourInADay yg diakses dari library.go

	// fmt.Println("minutesInHour = ", library.minutesInHour) //error
	// fmt.Println("secondInMinute = ", library.secondInMinute) //error

	fmt.Println("Name = ", library.Name)

	fmt.Println(library.DaysToMinutes(3))
	// fmt.Println(library.daysToHour(3)) //error
}