package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Pr() {
	if err := recover(); err != nil {
		fmt.Println("Recovered in f", err)
	}
}

func main() {
	a := 1 // line 1
	defer Pr()
	defer func() {
		s := a
		calc("1", a, calc("10", a, 1))
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err, s)
		}
	}()
	b := 2                               // line 2
	defer calc("1", a, calc("10", a, b)) // line 3
	a = 0                                // line 4
	panic(b)
	defer calc("2", a, calc("20", a, b)) // line 5
	b = 1                                // line 6
}
