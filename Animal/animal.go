package Animal

import "fmt"

type Animal struct {
	name  string
	tiles [][2]int
}

func NewDefault() Animal {
	return Animal{"Empty", make([][2]int, 0)}
}

func NewAnimal(name string, tiles [][2]int) Animal {
	return Animal{name, tiles}
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a *Animal) Print() {
	fmt.Printf(a.name)
}
