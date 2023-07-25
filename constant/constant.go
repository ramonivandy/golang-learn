package main

import "fmt"

func main() {
	/* 
		Welcome to constant!
	*/
	const firstName string = "Ramon"
	const lastName = "Setiawan"
	const val = 1000

	const (
		secondName = "Ramon"
		lastSecondName = "Ivandy"
		value = 2000
	)
	fmt.Println(firstName);
	fmt.Printf("%s %s \n", firstName, lastName)
	fmt.Println(val)

	fmt.Println(secondName)
	fmt.Println(lastSecondName)
	fmt.Println(value)
}