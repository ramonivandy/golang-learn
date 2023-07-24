package main

import "fmt"

func main()  {
	/* 
		String = kumpulan karakter
	*/
	fmt.Println("Ramon")
	fmt.Println("Ramon Ivandy")
	fmt.Println("Ramon Ivandy Setiawan")
	/* 
		function untuk string:
			1. len("string") = menghitung panjang karakter string
			2. "string"[number] = mengambil sebuah karakter pada "string" (index)
	*/
	fmt.Println(len("Ramon")) // return 3
	fmt.Println("Ramon Ivandy"[0]) // return 82 -> 82 adalah byte dari karakter R
	fmt.Println("Ramon Ivandy Setiawan"[1]) // 97 -> 97 adalah byte dari karakter a
}