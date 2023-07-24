package main

import "fmt"
func main() {
	/*
		Welcome to Integer!
		Tipe data integer (1): *punya nilai negatif
			1. int8 = -128 <> 127
			2. int16 = -32768 <> 32767
			3. int32 = -2 Miliar <> 2 Miliar
			4. int64 = lebih dari 2 Miliar
		Tipe data integer (2): *tidak punya nilai negatif
			uint = unsigned integer
			1. uint8 = 0 <> 255
			2. uint16 = 0 <> 65535
			3. uint32 = 0 <> 4 Miliar
			4. uint64 = 0 <> lebih dari 4 Miliar
		Tipe data float :
			1. float32 = 1.18x10^-38 <> 3.4x10^38
			2. float64 = 2.23x10^-308 <> 1.80x10^308
			3. complex64 (jarang digunakan)
			4. complex128 (jarang digunakan)
		"Alias" tipe data:
			1. byte = uint8
			2. rune = int32
			3. int = minimal int32/64 (menyesuaikan dengan OS)
			4. uint = minimal uint32/64 (menyesuaikan dengan Ps)
	*/

	fmt.Println("Satu = ", 1);
	fmt.Println("Dua = ", 2);
	fmt.Println("Tiga = ", 3.5);
}