package main

import "fmt"

func main() {
	ppls := []string{"HuXn", "John", "Jordan"}
	for index, value := range ppls {
		fmt.Println(index, value)
	}
}