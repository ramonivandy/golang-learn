package main

import "fmt"

func main()  {
	type NoKTP string
	type Married bool

	var noKtpRamon NoKTP = "32160909090909"
	var marriedStatus Married = false

	fmt.Println(noKtpRamon)
	fmt.Println(marriedStatus)
}