package main

import "fmt"

func main()  {
	name := "ramon"

	switch name {
	case "ramon":
		fmt.Println("Hello ramon")
		fmt.Println("Hello ramon 2")
	case "eko":
		fmt.Println("Hello eko")
		fmt.Println("Hello eko 2")
	default:
		fmt.Println("Hello no name")
		fmt.Println("Hello no name 2")
	}
	/* 
		switch with short expression
	*/
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("True!")
	case false:
		fmt.Println("False..")
	}

	/* 
		switch without condition
	*/
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama cukup panjang")
	case length > 10:
		fmt.Println("Nama lumayan panjang")
	case length > 20:
		fmt.Println("Nama kepanjaaaangan")
	default:
		fmt.Println("Nama udah bener")
	}
}