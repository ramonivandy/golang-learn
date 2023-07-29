package main

import "fmt"

/*
	semua variable di golang passing by value not reference.
	jika pengirim variable ke dalam function, yang dikirim adalah duplikasi value nya
*/

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia (address *Address) {
	address.Country = "Indonesia"
}

func main()  {
	/* 
		Golang do pass by value by default
	*/
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println("//=====")

	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("//=====")
	/* 
		Golang do pass by reference
		tambahan "&" di depan variable yang ingin di reference.
	*/
	address3 := &address1
	address3.City = "Cirebon"
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println("//=====")

	/* 
		Operator * di pointer
	*/
	*address3 = Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	/* 
		Function New
		membuat Address baru, tetapi dengan data yang kosong
	*/
	var address4 *Address = new(Address)
	fmt.Println(address4) // return &{ } (data kosong)
	address4.City = "Bali"
	fmt.Println(address4) // return &{"Bali"}

	var alamat = Address {
		City: "Subang",
		Province: "Jawa Barat",
		Country: "America",
	}
	fmt.Println(alamat)
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
	
}