package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
   runApplication(10)
}

func logging () {
	fmt.Println("Logging!")
}

/* 
	Defer func() digunakan untuk menjalankan function walaupun terdapat error, cocok untuk function logging.
*/
func runApplication (value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}