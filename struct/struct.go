package main

import "fmt"

/*
	Struct, untuk define dan dapat digunakan untuk menyimpan data. Lebih baik menggunakan struct dibanding array atau map.
*/
type Customer struct {
	Name, Address 	string
	Age 			int
}

func main()  {
	/* 
		Basic menggunakan struct
	*/
   var user Customer
   user.Name = "Ramon"
   user.Address = "Indonesia"
   user.Age = 25
   fmt.Println(user)

   /* 
	Struct literals
   */
   var user2 Customer = Customer {
		Name: "Ramon",
		Address: "Kedasih",
		Age: 30,
   }
   fmt.Println(user2)

   /* 
 	call struct function / struct method
	should add struct data first then call struct function 
   */
   user2.sayHi("Mochi")
}

/* 
	Function with struct / struct method
*/
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}