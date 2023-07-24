package main

import "fmt"

func main() {
	/*
		Variable
		Pada golang, variable hanya bisa menyimpan data yang sama (sejenis), string = string, int = int, bool = bool, dll.
	*/
	var name string // var => untuk deklarasi variable, name => nama variable, string => jenis data untuk variable
	name = "Ramon Ivandy"
	fmt.Println(name)
	name = "Ramon Ivandy Setiawan"
	fmt.Println(name)
	/* 
		Inisiasi langsung variable	
	*/
	var catto = "mochi"
	fmt.Println(catto)
	var mochiAge = 2
	fmt.Println(mochiAge)
	/* 
		Inisiasi langsung variable tanpa menggunakan "var"
	*/
	machuName := "Macha" 
	fmt.Println(machuName)
	machuAge := 3
	fmt.Println(machuAge)
	/* 
		Inisiasi variable bersamaan (lebih dari 2)
	*/
	var (
		firstName = "Ramon"
		lastName = "Setiawan"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}