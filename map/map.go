package main

import "fmt"

func main()  {
	person := map[string]string{
		"name": "ramon",
		"address": "cikarang",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "developer"
	fmt.Println(person["title"])

	 var book map[string]string = make(map[string]string)

	 book["title"] = "Belajar Go-Lang!"
	 book["author"] = "Ramon"
	 book["this"] = "wrong"
	 fmt.Println(book)
	 delete(book, "this")
	 fmt.Println(book)
}