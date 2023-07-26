package main

import "fmt"

func main()  {

   goodbye := sayGoodbye
   fmt.Println(goodbye("Ramon"))
}

func sayGoodbye(name string) string {
   return "Goodbye... " + name
}