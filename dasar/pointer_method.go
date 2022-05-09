package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	foo := Man{"Bayu"}
	foo.married()

	fmt.Println(foo.Name)
}
