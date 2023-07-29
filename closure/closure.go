package main

import "fmt"

func main()  {
	name := "Ramon"
	counter := 0

	/* 
		melakukan perubahan pada variable diluar anonim func 
	*/
	increment := func() {
		name = "Ivandy"
		counter++
	}

	increment()

	fmt.Println(name)
	fmt.Println(counter)
}

