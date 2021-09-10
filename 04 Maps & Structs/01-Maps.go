package main

import "fmt"

func main() {
	userData := map[string]int {
		"huxn": 17,
		"alex": 18,
		"john": 27,
	}
	fmt.Println(userData)
	userData["jordan"] = 15
	fmt.Println(userData["huxn"])
	fmt.Println(userData["alex"])
	fmt.Println(userData["john"])

	for prop, value := range userData {
		fmt.Println(prop, value)
	}

	fmt.Println(userData)
	delete(userData, "john")
	fmt.Println(userData)

}