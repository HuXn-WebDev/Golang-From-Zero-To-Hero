package main

import "fmt"

func main() {
	// num := []int{10,20,30,40,50}
	// fmt.Println(num)

	// fmt.Println(num[:])
	// fmt.Println(num[1:])
	// fmt.Println(num[1:3])

	// num = append(num, 60, 70, 80, 90)
	// fmt.Println(num)

	num := make([]int, 5, 20)
	fmt.Println(num)

	num[2] = 5
	fmt.Println(num)
	num[6] = 30
	fmt.Println(num)
}