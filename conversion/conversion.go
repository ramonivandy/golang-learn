package main

import "fmt"

func main()  {
	// conversi nilai integer ke integer 
	const nilai32 int32 = 127
	const nilai64 int64 = int64(nilai32)
	const nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//conversi nilai byte menjadi string
	var name = "ramon"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}