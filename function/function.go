package main

import "fmt"

func main()  {
	slice := make([]string, 3, 5)
	slice[0] = "ramon"
	slice[1] = "eko"
	slice[2] = "budi"
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		sayHello(slice[i], i)
	}

	/* func return data */
	fmt.Println(sumNumber(1, 2))

	/* func return multiple data */
	firstName, lastName := getFullname()
	fmt.Println(firstName, lastName)

	/* 
		underscore `_` => digunakan untuk ignore return suatu value dari multiple return func
	*/
	firstName2, _ := getFullname()
	fmt.Println(firstName2)

	/* 
		func return named value 
	*/
	firstName3, middlename3, lastname3 := getCompleteName()
	fmt.Println(firstName3, middlename3, lastname3)
}
/* 
	go function with parameter
*/
func sayHello(name string, index int) {
	fmt.Println("Hello!", name, "with key", index)
}
/* 
	go function return value
*/
func sumNumber(number1 int, number2 int) int {
	result := number1 + number2
	return result
}
/* 
	go func return multiple value
*/
func getFullname() (string, string) {
	firstName := "Ramon"
	lastName := "Ivandy"

	return firstName, lastName
}
/* 
	go func named return values
*/
func getCompleteName() (firstname string, middlename string, lastname string) {
	firstname = "Ramon"
	middlename = "Ivandy"
	lastname = "Setiawan"

	return
}