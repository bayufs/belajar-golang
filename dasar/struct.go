package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi nama saya ", customer.Name, "Suami dari ", name)
}

func main() {

	var bayu Customer

	bayu.Name = "Bayu"
	bayu.Address = "ckr"
	bayu.Age = 29

	fmt.Println(bayu.Name)
	fmt.Println(bayu.Address)
	fmt.Println(bayu.Age)

	lia := &Customer{
		Name:    "lia",
		Address: "Kapuk",
		Age:     28,
	}
	fmt.Println(lia.Name)
	fmt.Println(lia.Address)
	fmt.Println(lia.Age)

	bayu.sayHi("Lia")

}
