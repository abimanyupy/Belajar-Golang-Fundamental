package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "E:/DATA/CODE/GOLANG/ENIGMA/GOLANG FUNDAMENTAL/DATA/"
var fileName = "example.txt"
var filePath = path + fileName

func main() {
	fmt.Println("=======================================")
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer dir.Close() // Defer closing the directory

	// Read directory entries
	entries, err := dir.ReadDir(0) // Read all entries
	if err != nil {
		panic(err)
	}

	// Display file names
	fmt.Println("Daftar File:")
	for _, entry := range entries {
		if !entry.IsDir() {
			// Extract filename without path
			filename := entry.Name()
			for i := len(filename) - 1; i >= 0; i-- {
				if filename[i] == '/' {
					filename = filename[i+1:]
					break
				}
			}
			fmt.Println(filename)
		}
	}
	fmt.Println("=======================================")

	fmt.Print("1. Membuat file \n")
	fmt.Print("2. Menulis ke file \n")
	fmt.Print("3. Memilih file yang ingin dibaca \n")
	fmt.Print("4. Tampilkan semua file \n")
	fmt.Print("5. Hapus file \n")
	fmt.Print("Pilih Menu : ")

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	switch input {
	case "1":
		createFile()
	case "2":
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Masukkan nama tujuan File: ")
		scanner.Scan()
		fileName = scanner.Text() + ".txt"
		filePath = path + fileName

		fmt.Print("Masukkan teks : ")
		scanner.Scan()
		input = scanner.Text()
		writeFile(input)
	case "3":
		readFile() //untuk membaca file .txt
	case "4":
		dir, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer dir.Close()
		input, err := dir.ReadDir(0)
		if err != nil {
			panic(err)
		}
		tampilFile(input)
	case "5":
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Masukkan nama file yang ingin di hapus: ")
		scanner.Scan()
		fileName = scanner.Text() + ".txt"
		filePath = path + fileName
		hapusFile()
	default:
		fmt.Println("Input tidak valid")
	}

	// if input == "1" {
	// 	createFile()
	// } else if input == "2" {

	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	fmt.Print("Masukkan nama tujuan File: ")
	// 	scanner.Scan()
	// 	fileName = scanner.Text() + ".txt"
	// 	filePath = path + fileName

	// 	fmt.Print("Masukkan teks : ")
	// 	scanner.Scan()
	// 	input = scanner.Text()
	// 	writeFile(input)
	// } else if input == "3" {
	// 	readFile() //untuk membaca file .txt
	// } else if input == "4" {
	// 	dir, err := os.Open(path)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer dir.Close()
	// 	input, err := dir.ReadDir(0)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	tampilFile(input)
	// } else {
	// 	fmt.Println("Input tidak valid")
	// }

}

func createFile() { //fungsi untuk membuat file .txt
	var filePathh string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama file yang ingin dibuat: ")
	scanner.Scan()
	fileNamee := scanner.Text() + ".txt" // Get the file name from user input

	if scanner.Text() == "" {
		fmt.Println("File name cannot be empty")
		return
	}

	filePathh = path + fileNamee

	_, err := os.Stat(filePathh) //ada 2 return, 1.informasi ttg path yg dicar 2. error jika ada
	if os.IsNotExist(err) {      //jika path yg dicari tidak ada
		file, err := os.Create(filePathh)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close() //untuk menutup file

		fmt.Println("File", fileNamee, "Berhasil dibuat")
	}
}

func writeFile(input string) {

	// file, err := os.OpenFile(filePath, os.O_RDWR, 0666) //ini akan menimpa hasil inputan awalnya
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666) //untuk open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// writer.WriteString(input) //akan menerima argumen string yg string tsb akan ditampung pada suatu tempat
	// writer.WriteString("\n") //menuliskan string \n untuk enter baris

	writer.WriteString(input + "\n") //menggabungkan input string dan \n
	writer.Flush()                   //yg akan menuliskan semua string yg sudah ditampung tadi ke dalam file
}

func readFile() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama File : ")
	scanner.Scan()
	input = scanner.Text()

	filePath := path + input + ".txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scannerr := bufio.NewScanner(file)
	for scannerr.Scan() {
		fmt.Println(scannerr.Text())
	}
}

func tampilFile(entries []os.DirEntry) {
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}

func hapusFile() {
	var err = os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file", fileName, "berhasil di delete")
}
