package main

import "fmt"

func main() {
	address := 10
	fmt.Println(address)
	fmt.Println(&address)
	fmt.Printf("%T\n", address)
	fmt.Printf("%T\n", &address)

	value := &address
	fmt.Println(*&value)
}