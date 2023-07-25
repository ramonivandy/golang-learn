package main

import "fmt"

func main()  {
   fmt.Println("Hello world")

   //operasi penambahan
   var result = 10 + 10
   fmt.Println(result)

   //operasi perkalian
   var (
		a = 10
		b = 10
   )
   fmt.Println(a * b)

   /* 
 	Augmented Assignments  
   */
   var i = 20
   i += 10 //mempersingkat operasi math
   fmt.Println(i)

   /* 
 	Unary Operation  
   */
   var j = 10
   j++
   fmt.Println(j)
}