package main

import "fmt"

func main()  {
   fmt.Println("Hello world")

   fmt.Println(factorialLoop(5))

   fmt.Println(factorialRecursive(5))
}

/* 
	Recursive..

	Jadi, saat value = 5, maka akan melakukan 
		return => "5 * factorialRecursive(4)"
	lalu, lanjut melakukan 
		return => "4 * factorialRecursive(3)"
	...sampai di value = 1
		factorialRecursive(1) akan melakukan return integer 1
	dari sini perkalian dimulai
		return 2 * factorialRecursive(1) berubah menjadi 2 * 1 = 2
	dan terus berlanjut...
		return 3 * factorialRecursive(2) berubah menjadi 3 * 2 = 6
		return 4 * factorialRecursive(3) berubah menjadi 4 * 6 = 24
		return 5 * factorialRecursive(4) berubah menjadi 5 * 24 = 120
	Wow *dogesmirk 
*/
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}