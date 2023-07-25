package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Berubah"
	fmt.Println(slice1)

	/* 
		Slice function =>
		1. len(slice) = panjang slice
		2. cap(slice) = kapasitas slice
		3. append(slice, data) = membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas penuh membuat array baru
		4. make([]TypeData, length, capacity) = membuat slice baru
		5. copy(destination, source) = menyalin slice dari source ke destination
	*/

	//append
	var slice2 = months[10:]
	fmt.Println(slice2)
	
	var slice3 = append(slice2, "Ramon")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)

	//make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ramon"
	newSlice[1] = "Setiawan"

	fmt.Println(newSlice)

	//copy
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	//warning
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}