package main

import "fmt"

type Blacklist func(string) bool

func main()  {
	/* 
		Anonymous function
		pembuatan function langsung tanpa harus deklarasi.
	*/
	blacklist := func(name string) bool {
		return name == "admin"
	}
	
	registerUser("admin", blacklist)
	registerUser("ramon", blacklist)
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome!", name)
	}
}