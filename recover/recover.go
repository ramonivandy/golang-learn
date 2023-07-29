package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
   runApp(true)
   fmt.Println("Make sure app not stop")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error!") // terdapat panic, harus melakukan recover di defer func. Jika tidak maka aplikasi akan berhenti
	}
	fmt.Println("Aplikasi berjalan")
}

/* 
		recover harus disimpan diluar function yang terdapat panic.
		biasanya recover disimpan pada defer dari function yang terdapat panic
*/
func endApp()  {
	message := recover() /* panic terdapat di func runApp dengan defer endApp, lakukan cek dengan recover apakah panic dieksekusi atau tidak */
	if message != nil {
		fmt.Println("Terdapat error dengan msg =", message)
	}
	fmt.Println("Aplikasi selesai")
}