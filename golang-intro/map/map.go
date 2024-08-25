package main

import "fmt"

func main() {
	user := map[string]string{ //map[tipe data key] tipe data value
		"username": "abimanyu", //username adalah key dan abimanyu adalah value
		"email": 	"abimanyu@mail.com",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

	fmt.Println("==========================================")

	var skor = make(map[string]int)
	fmt.Println(skor) 

	skor["java"] = 90
	skor["react"] = 80
	skor["kotlin"] = 88

	fmt.Println("skor :", skor)

	fmt.Println("nilai java :", skor["java"])
	fmt.Println("nilai react :", skor["react"])
	fmt.Println("nilai golang :", skor["golang"])

	fmt.Println("==========================================")

	skor["java"] = 97

	fmt.Println("skor :", skor)

	fmt.Println("==========================================")

	delete(skor, "kotlin")  //nama map dan key
	fmt.Println("skor :", skor)

	fmt.Println("==========================================")

	names := map[int]string{
		1: "abimanyu",
		2: "jak",
		3: "jun",
		4: "lam",
	}
	for key, value := range names {
		fmt.Println(key, ":", value)
	}

	for key := range names {
		fmt.Println("key :", key)
	}

	for _, value := range names {
		fmt.Println("value :", value)
	}

}