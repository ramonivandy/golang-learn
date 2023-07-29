package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
   runApplication(false)
}

func runApplication(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error!")
	}
	fmt.Println("Aplikasi jalan")
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}