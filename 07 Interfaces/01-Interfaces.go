package main

import "fmt"

type human interface {
	printInfo()
}

type person struct {
	name string
}

type programmer struct {
	person
	isProgrammer bool
}

func (p person) printInfo() {
	fmt.Println("#Person: Name:", p.name)
}

func (cp programmer) printInfo() {
	fmt.Println("#Programmer: Name:", cp.name, "Is Programmer", cp.isProgrammer)
}

func show(h human) {
	h.printInfo()
}

func main() {
	huxn := person {
		name: "HuXn",
	}

	john := programmer {
		person: person {
			name: "John",
		},
		isProgrammer: true,
	}

	var p1, p2 human
	p1 = huxn
	p2 = john
	p1.printInfo()
	p2.printInfo()
	fmt.Println("-------------------")
	show(p1)
	show(p2)
	show(huxn)
	show(john)
}