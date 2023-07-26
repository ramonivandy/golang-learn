package main

import "fmt"

func main()  {
   fmt.Println("Hello world")
	/* 
		break menghentikan perulangan
	*/
	for i := 0; i < 10; i++ {
		fmt.Println("index =", i)
		if(i == 5) {
			break
		}
	}

	/* 
		continue melakukan skip pada perulangan n
	*/
	for i := 0; i < 20; i++ {
		if(i % 2 == 1){
			continue
		}

		fmt.Println("genap =>", i)
	}

}