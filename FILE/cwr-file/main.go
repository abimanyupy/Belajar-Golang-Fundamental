package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var path = "E:/DATA/CODE/GOLANG/ENIGMA/GOLANG FUNDAMENTAL/DATA/" 
// var fileName = "example.txt"
// var filePath = path + fileName

// func main() {
// 	var input string
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Masukkan nama : ")
// 	scanner.Scan()
// 	input = scanner.Text()

// 	writeFile(input) //untuk menulis ke file .txt

// 	readFile() //untuk membaca file .txt
// }

// func createFile() { //fungsi untuk membuat file .txt
// 	_, err := os.Stat(filePath) //ada 2 return, 1.informasi ttg path yg dicar 2. error jika ada
// 	if os.IsNotExist(err) {     //jika path yg dicari tidak ada
// 		file, err := os.Create(filePath)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		defer file.Close() //untuk menutup file

// 		fmt.Println("File", fileName, "Berhasil dibuat")
// 	}
// }

// func writeFile(input string) {
// 	// file, err := os.OpenFile(filePath, os.O_RDWR, 0666) //ini akan menimpa hasil inputan awalnya
// 	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666) //untuk open file
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	// writer.WriteString(input) //akan menerima argumen string yg string tsb akan ditampung pada suatu tempat
// 	// writer.WriteString("\n") //menuliskan string \n untuk enter baris

// 	writer.WriteString(input + "\n") //menggabungkan input string dan \n
// 	writer.Flush()                   //yg akan menuliskan semua string yg sudah ditampung tadi ke dalam file
// }

// func readFile() {
// 	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666) //untuk open file
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }
