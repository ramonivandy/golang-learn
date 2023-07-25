package main

import "fmt"

func main()  {
   fmt.Println("Hello world")

   var name1 = "Ramon"
   var name2 = "Amon"
   var result bool = name1 == name2
   fmt.Println(result)

   var num1 = 100
   var num2 = 300

   fmt.Println(num1 > num2)
   fmt.Println(num1 < num2)
   fmt.Println(num1 == num2)
   fmt.Println(num1 != num2)

}