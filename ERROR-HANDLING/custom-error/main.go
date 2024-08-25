package main

import (
	// "errors"
	"fmt"
)

type Item struct {
	id    int
	name  string
	price int
}

type ErrorItemNotFound struct { //cara utk membuat custom error //struct ini harus memiliki method yg mengimplementasikan interface error yg dimana method tsb adalah Error() string
	id int
}

func (i ErrorItemNotFound) Error() string { //dimana methodnya Error() string
	return fmt.Sprintf("item dengan id: %d tidak ditemukan", i.id)
}

var items = []Item{
	{1, "kursi", 200000},
	{2, "meja", 100000},
	{3, "lemari", 220000},
}

func getItemById(id int) (Item, error) {
	for _, item := range items {
		if item.id == id {
			return item, nil
		}
	}
	return Item{}, &ErrorItemNotFound{id}
}

func main() {
	result, err := getItemById(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
