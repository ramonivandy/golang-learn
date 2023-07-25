package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
	var names [3]string
	names[0] = "Ramon"
	names[1] = "Ivandy"
	names[2] = "Setiawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//======
	var values = [3]int {
		90,
		95,
		80,
	}
	fmt.Println(values)

	/* 
		Array function di Go lang
		len(array) = mendapatkan panjang array
		array[index] = mendapatkan data di posisi index
		array[index] = value > mengubah nilai value pada index
	*/

	fmt.Println(len(names))
	fmt.Println(len(values))

	//jika sudah membuat array, walaupun datanya kosong panjang array tetap sesuai yang sudah dibuat
	var jumlah [10]string
	var jumlah2 = [10]string{}
	fmt.Println(len(jumlah))
	fmt.Println(len(jumlah2))

}