package qcode

import "fmt"

// type Address struct {
// 	City, State string
// }

// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

type Person struct {
	Info
	Address
}

type Info struct {
	name string
	age  int
}

type Address struct {
	add string
}

func (p Person) PersonInfo() {
	fmt.Printf("Name is %s, %d year old. Live in %s", p.name, p.age, p.Address)
}

func PersonInfoFunc() {
	p := Person{Info: Info{name: "john", age: 22}, Address: Address{"LA in unit state."}}
	p.PersonInfo()
}
