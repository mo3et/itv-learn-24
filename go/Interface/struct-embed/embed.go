package main

// struct 内嵌 Interface

// 例如 A和B 都可以实现 Itf接口

type Embed interface {
	define(string)
}

type Humans struct {
	Embed []Humans
	name  string
}

func (h *Humans) define(job string) {
}

func (h *Humans) drive() {
}

type child struct {
	name string
}

func (c *child) define() {
}
func (c *child) run() {}

func main() {
}
