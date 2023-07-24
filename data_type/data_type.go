package main

import "fmt"

func main() {
	fmt.Println("Welcome to Data Type in Go!")
	/* table of go numbers */
	/*
		uint8	0 ↔ 255
		uint16	0 ↔ 65535
		uint32	0 ↔ 4294967295
		uint64	0 ↔ 18446744073709551615
		uint	sama dengan uint32 atau uint64 (tergantung nilai)
		byte	sama dengan uint8
		int8	-128 ↔ 127
		int16	-32768 ↔ 32767
		int32	-2147483648 ↔ 2147483647
		int64	-9223372036854775808 ↔ 9223372036854775807
		int	sama dengan int32 atau int64 (tergantung nilai)
		rune	sama dengan int32
	*/

	//positive and negative numbers
	var positive uint8 = 2
	negative := -1234567
	fmt.Printf("Nomor Positif: %d, negatif: %d \n", positive, negative)

	//decimal number
	decimal := 2.622
	// %f for decimal value, %.2f for max 2 number after decimal point. Change the number to decimal want to showed
	fmt.Printf("Decimal Number: %f \n", decimal)
	fmt.Printf("Decimal Number with Max 2 digit: %.2f \n", decimal)

	//bool value
	boolValue := false
	boolValue2 := true
	// %t for boolean values
	fmt.Printf("IsFalse: %t, isTrue: %t \n", boolValue, boolValue2)

	//string type
	var string = "hello world"
	fmt.Println(string)

	//string with backtick
	var string_with_backtick = `
		Hello
		From 
		Backtick 
		Variable
	`
	fmt.Println(string_with_backtick)

	// nil value
	/*
		nil != zero value;
		string zero value = ""
		integer zero value = 0
		decimal zero value = 0.0
		boolean zero value = false

		nil can be used to this type data :
		- pointer
		- function data type
		- slice
		- map
		- channel
		- empty interface ( interface{} )
	*/
}
