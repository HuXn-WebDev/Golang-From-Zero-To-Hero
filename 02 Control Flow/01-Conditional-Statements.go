package main

import "fmt"

func main() {
	password := "12345678"
	if len(password) > 7 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}
}