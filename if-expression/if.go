package main

import "fmt"

func main()  {
	name := "ramonivandy"	
	if name == "ramonivandy" {
		fmt.Println("Hello ramon")
	} else if name == "eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hello no name")
	}

	/* 
		if short statement 
		menarik! Namun sepertinya agak sulit dibaca kalau tidak teliti :/
	*/
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah oke")
	}
}