package main

import "fmt"

/*
	Pointer bertujuan untuk melakukan hemat memory, dikarenakan apabila tidak menggunakan pointer maka data akan dilakukan diplikasi terus menerus
	akan berbahaya apabila struct yang dibikin sudah memiliki banyak field
*/
type Man struct {
	Name string
}

func (man *Man) Married () {
	man.Name = "Mr. " + man.Name
}

func main()  {
	ramon := Man{"Ramon"}
	ramon.Married()
   fmt.Println(ramon)
}