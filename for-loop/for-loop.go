package main

import "fmt"

func main()  {
	counter := 1
	/* 
		more looks like while expression, not for expression...
	*/
	for counter <= 10 {
		fmt.Println("counter ke", counter)
		counter++
	}
   /* 
		for dengan statement
		this is the for i knew!
   */
	for i := 1; i <= 10; i++ {
		fmt.Println("perularang i ke -", i)
	}
	/* 
		for range
	*/
	slice := []string{"ramon","ivandy","setiawan"}
	// normal loop
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	/* 
		for range, its more like this in javascript => for (const val in array) {}
	*/
	for idx, val := range slice {
		fmt.Println("index ke -", idx, "dengan value =", val)
	}
	/* 
		for range in map
	*/
	person := make(map[string]string)
	person["name"] = "ramon"
	person["title"] = "developer"
	person["birthPlace"] = "bandung"
	person["birthDate"] = "09/05/1998"
	for key, val := range person {
		fmt.Println("key", key, "dengan value =", val)
	}
}