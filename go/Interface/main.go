package main

import (
	"fmt"

	codeq "github.com/mo3et/itv-learn-24/go/Interface/code-Q"
)

type (
	Dog    struct{}
	Rabbit struct {
		name   string
		number int
	}
)

func (d *Dog) Moved() string {
	return "Dog Moved."
}

func (d Dog) Speak() {
	fmt.Println("Fuxk you")
}

func (r *Rabbit) Moved() string {
	// name := r.name
	r.name = "no true name"

	return r.name
}

func (r *Rabbit) Speak() {
	fmt.Println("nope")
}

type Animal interface {
	Moved() string
	Speak()
}

func main() {
	var _ Animal = (*Rabbit)(nil)
	var _ Animal = (*Dog)(nil)

	// var d Animal = &Dog{}
	// rimpl := &Rabbit{}
	// var r *Rabbit = &Rabbit{}
	// var a Animal = &Rabbit{}
	// d.Moved()
	// d.Speak()
	// r.Moved()
	// a.Moved()
	// rimpl.Moved()
	codeq.Q1()
	codeq.Q2()
	codeq.Q3()
	codeq.Q4()
	// codeq.Q5()
	// codeq.Q6()
}
