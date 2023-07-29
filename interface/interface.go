package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main()  {
	var ramon Person
	ramon.Name = "Ramoon"
   	SayHello(ramon)

	var cat Animal
	cat.Name = "Mochii"
	SayHello(cat)
}