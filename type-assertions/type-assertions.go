package main

import "fmt"

func random() interface{} {
	return false
}

func main()  {
	var result interface{} = random()
	// var resultString string = result.(string)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}