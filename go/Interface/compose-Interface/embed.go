package main

import "fmt"

// struct 内嵌 Interface slice
// 组合接口

type Ningen interface {
	Define(string)
}

type Humans struct {
	Ningens []Ningen
	name    string
}

func (h *Humans) Define(name string) {
	fmt.Println("Human name is", h.name)
}

type Child struct {
	name string
}

func (c *Child) Define(name string) {
	fmt.Println("child name is", c.name)
}

func main() {
	human1 := &Humans{
		Ningens: []Ningen{
			&Humans{
				name: "Jason",
			},
			&Child{
				name: "lina",
			},
		},
	}
	for _, human := range human1.Ningens {
		human.Define(human1.name)
	}
}
