package main

import (
	"fmt"

	animal "github.com/discozoo/Animal"
)

func main() {
	fmt.Print("hello world \n")
	a := animal.NewDefault()
	a.SetName("Sheep")
	fmt.Print(a)

}
