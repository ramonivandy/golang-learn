package main

import "fmt"

type Filter func(string) string

func main()  {
   fmt.Println(sayHello("Ramon", spamFilter))
   fmt.Println(sayHello("Mochi", spamFilter))
   fmt.Println(sayHello("Anjing", spamFilter))
}

func sayHello(name string, filter Filter) string {
	namedFiltered := filter(name)
	return "Hello " + namedFiltered
}
 
func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}