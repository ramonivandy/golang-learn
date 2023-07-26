package main

import "fmt"

func main()  {
   fmt.Println(sumAll(10,20,30,40,50))

   slices := []int{20,30,40,50,60}
   fmt.Println(sumAll(slices...))
}

func sumAll(numbers ...int) int {
	count := 0;
	// for i := 0; i < len(numbers); i++ {
	// 	count += numbers[i]
	// } 
	for _, val := range numbers {
		count += val
	}
	return count
}