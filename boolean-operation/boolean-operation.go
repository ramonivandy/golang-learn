package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
   /* 
	Boolean:
	&& => Dan
	|| => Atau
	! = sebaliknya
   */

   /* 
	&& >
	true && true = true
	true && false = false
	false && true = false
	false && false = false
   */

    /* 
	|| >
	true || true = true
	true || false = true
	false || true = true
	false || false = false
   */

   /* 
	! >
	!true = false
	!false = true
   */

	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80
	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}