package main

import "fmt"


type Books struct {
	title string
	author string
}

func main(){
	var book1 Books
	book1.title = "Go 语言入门"
	book1.author = "mars.hao"
	fmt.Println(book1)
	fmt.Println(book1.title)
}
