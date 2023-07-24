package main

import "fmt"

func main() {
	//declare single variable with type
	var firstName string = "john"

	//declare variable and type, then set it later.
	var lastName string
	notLastName := "notLast Name"
	lastName = "doe"
	fmt.Printf("halo %s %s %s \n", firstName, lastName, notLastName)

	//declare multiple variables
	var umur, tempat, tanggal string = "25", "Bandung", "9 Mei 1998"
	fmt.Printf("Umur %s, lahir di %s pada tanggal %s \n", umur, tempat, tanggal)

	//declare variable without type
	var nonType = "no type"
	simpleNoType := "no type"
	fmt.Printf("%s %s \n", nonType, simpleNoType)

	//trash variable
	//any value that inserted here will removed permanently
	_ = "var sampah"

	//pointer variable
	pointerVar := new(string)

	// fmt.Printf(pointerVar) // if we do this, go will print memory address not the value of variable
	fmt.Println(*pointerVar) // add asterisk * to show variable value from pointer

}
