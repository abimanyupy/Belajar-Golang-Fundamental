package library

const secondInMinute int = 60
const minutesInHour int = 60 //ini contoh UNEXPORTED
const HourInADay int = 24  //membuaat variaable ini EXPORTED dan hanya ini yg bisa diakses di package berbeda

func daysToHours(day int) int { //merupakan function UNEXPORTED
	return day * HourInADay
}

func DaysToMinutes(day int) int { //merupakan function EXPORTED
	return day * HourInADay * minutesInHour
}

